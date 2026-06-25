
/**
 * Samples for L3Networks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/L3Networks_Get.json
     */
    /**
     * Sample code: Get L3 network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getL3Network(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l3Networks().getByResourceGroupWithResponse("resourceGroupName", "l3NetworkName",
            com.azure.core.util.Context.NONE);
    }
}
