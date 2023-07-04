/** Samples for StorageAppliances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/StorageAppliances_Get.json
     */
    /**
     * Sample code: Get storage appliance.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getStorageAppliance(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .storageAppliances()
            .getByResourceGroupWithResponse(
                "resourceGroupName", "storageApplianceName", com.azure.core.util.Context.NONE);
    }
}
