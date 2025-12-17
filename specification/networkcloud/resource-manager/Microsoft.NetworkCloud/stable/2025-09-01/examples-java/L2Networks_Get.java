
/**
 * Samples for L2Networks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/L2Networks_Get.json
     */
    /**
     * Sample code: Get L2 network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getL2Network(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l2Networks().getByResourceGroupWithResponse("resourceGroupName", "l2NetworkName",
            com.azure.core.util.Context.NONE);
    }
}
