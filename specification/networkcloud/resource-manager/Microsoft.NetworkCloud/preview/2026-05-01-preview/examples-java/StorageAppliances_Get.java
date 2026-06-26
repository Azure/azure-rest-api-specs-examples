
/**
 * Samples for StorageAppliances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/StorageAppliances_Get.json
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
