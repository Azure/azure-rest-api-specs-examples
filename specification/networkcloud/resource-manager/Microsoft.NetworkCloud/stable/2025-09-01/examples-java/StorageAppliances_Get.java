
/**
 * Samples for StorageAppliances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * StorageAppliances_Get.json
     */
    /**
     * Sample code: Get storage appliance.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getStorageAppliance(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.storageAppliances().getByResourceGroupWithResponse("resourceGroupName", "storageApplianceName",
            com.azure.core.util.Context.NONE);
    }
}
