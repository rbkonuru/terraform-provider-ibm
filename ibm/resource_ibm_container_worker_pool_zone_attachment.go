package ibm

import (
	"fmt"
	"strings"
	"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMContainerWorkerPoolZoneAttachment() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMContainerWorkerPoolZoneAttachmentCreate,
		Read:     resourceIBMContainerWorkerPoolZoneAttachmentRead,
		Update:   resourceIBMContainerWorkerPoolZoneAttachmentUpdate,
		Delete:   resourceIBMContainerWorkerPoolZoneAttachmentDelete,
		Exists:   resourceIBMContainerWorkerPoolZoneAttachmentExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
			Update: schema.DefaultTimeout(90 * time.Minute),
			Delete: schema.DefaultTimeout(90 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cluster": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"worker_pool": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"private_vlan_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"public_vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"worker_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceIBMContainerWorkerPoolZoneAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	zone := d.Get("zone").(string)
	workerPoolZoneNetwork := v1.WorkerPoolZoneNetwork{
		PrivateVLAN: d.Get("private_vlan_id").(string),
	}

	if v, ok := d.GetOk("public_vlan_id"); ok {
		workerPoolZoneNetwork.PublicVLAN = v.(string)
	}

	workerPoolZone := v1.WorkerPoolZone{
		ID: zone,
		WorkerPoolZoneNetwork: workerPoolZoneNetwork,
	}

	cluster := d.Get("cluster").(string)
	workerPool := d.Get("worker_pool").(string)

	workerPoolsAPI := csClient.WorkerPools()

	err = workerPoolsAPI.AddZone(cluster, workerPool, workerPoolZone)
	if err != nil {
		return err
	}
	_, err = WaitForWorkerZoneNormal(cluster, workerPool, zone, meta, d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return fmt.Errorf(
			"Error waiting for workers of worker pool (%s) of cluster (%s) to become ready: %s", workerPool, cluster, err)
	}
	d.SetId(fmt.Sprintf("%s/%s/%s", cluster, workerPool, zone))

	return resourceIBMContainerWorkerPoolZoneAttachmentRead(d, meta)

}

func resourceIBMContainerWorkerPoolZoneAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPool := parts[1]
	zoneName := parts[2]

	workerPoolsAPI := csClient.WorkerPools()

	workerPoolRes, err := workerPoolsAPI.GetWorkerPool(cluster, workerPool)
	if err != nil {
		return err
	}
	zones := workerPoolRes.Zones

	for _, zone := range zones {
		if zone.ID == zoneName {
			d.Set("public_vlan_id", zone.PublicVLAN)
			d.Set("private_vlan_id", zone.PrivateVLAN)
			d.Set("worker_count", zone.WorkerCount)
			d.Set("zone", zone.ID)
			d.Set("cluster", cluster)
			d.Set("worker_pool", workerPool)
			break
		}
	}

	return nil
}

func resourceIBMContainerWorkerPoolZoneAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	workerPoolsAPI := csClient.WorkerPools()

	if d.HasChange("private_vlan_id") || d.HasChange("public_vlan_id") {
		parts, err := idParts(d.Id())
		if err != nil {
			return err
		}
		cluster := parts[0]
		workerPool := parts[1]
		zone := parts[2]
		err = workerPoolsAPI.UpdateZoneNetwork(cluster, zone, workerPool, d.Get("private_vlan_id").(string), d.Get("public_vlan_id").(string))
		if err != nil {
			return err
		}
	}

	return resourceIBMContainerWorkerPoolZoneAttachmentRead(d, meta)
}

func resourceIBMContainerWorkerPoolZoneAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPool := parts[1]
	zone := parts[2]

	workerPoolsAPI := csClient.WorkerPools()
	err = workerPoolsAPI.RemoveZone(cluster, zone, workerPool)
	if err != nil {
		return err
	}
	_, err = WaitForWorkerZoneDeleted(cluster, workerPool, zone, meta, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return fmt.Errorf(
			"Error waiting for deleting workers of worker pool (%s) of cluster (%s):  %s", workerPool, cluster, err)
	}

	return nil
}

func resourceIBMContainerWorkerPoolZoneAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	cluster := parts[0]
	workerPoolID := parts[1]
	zoneID := parts[2]

	workerPoolsAPI := csClient.WorkerPools()

	workerPool, err := workerPoolsAPI.GetWorkerPool(cluster, workerPoolID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	zones := workerPool.Zones
	var zone v1.WorkerPoolZoneResponse
	for _, z := range zones {
		if z.ID == zoneID {
			zone = z
		}
	}
	return zone.ID == zoneID, nil
}

func WaitForWorkerZoneNormal(clusterNameOrID, workerPoolNameOrID, zone string, meta interface{}, timeout time.Duration) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", workerProvisioning},
		Target:     []string{workerNormal},
		Refresh:    workerPoolZoneStateRefreshFunc(csClient.Workers(), clusterNameOrID, workerPoolNameOrID, zone),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func workerPoolZoneStateRefreshFunc(client v1.Workers, instanceID, workerPoolNameOrID, zone string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		workerFields, err := client.ListByWorkerPool(instanceID, workerPoolNameOrID, false)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		//Done worker has two fields State and Status , so check for those 2
		for _, e := range workerFields {
			if e.Location == zone {
				if strings.Contains(e.KubeVersion, "pending") || strings.Compare(e.State, workerNormal) != 0 || strings.Compare(e.Status, workerReadyState) != 0 {
					if strings.Compare(e.State, "deleted") != 0 {
						return workerFields, workerProvisioning, nil
					}
				}
			}
		}
		return workerFields, workerNormal, nil
	}
}

func WaitForWorkerZoneDeleted(clusterNameOrID, workerPoolNameOrID, zone string, meta interface{}, timeout time.Duration) (interface{}, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return nil, err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{workerDeleteState},
		Refresh:    workerPoolZoneDeleteStateRefreshFunc(csClient.Workers(), clusterNameOrID, workerPoolNameOrID, zone),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func workerPoolZoneDeleteStateRefreshFunc(client v1.Workers, instanceID, workerPoolNameOrID, zone string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		workerFields, err := client.ListByWorkerPool(instanceID, workerPoolNameOrID, true)
		if err != nil {
			return nil, "", fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		//Done worker has two fields State and Status , so check for those 2
		for _, e := range workerFields {
			if e.Location == zone {
				if strings.Compare(e.State, "deleted") != 0 {
					return workerFields, "deleting", nil
				}
			}
		}
		return workerFields, workerDeleteState, nil
	}
}