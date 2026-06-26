
/**
 * Samples for Racks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Racks_Get.json
     */
    /**
     * Sample code: Get rack.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getRack(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.racks().getByResourceGroupWithResponse("resourceGroupName", "rackName",
            com.azure.core.util.Context.NONE);
    }
}
