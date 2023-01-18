/** Samples for VirtualNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_Get.json
     */
    /**
     * Sample code: VirtualNetworks_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualNetworksGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualNetworks()
            .getWithResponse(
                "resourceGroupName", "{labName}", "{virtualNetworkName}", null, com.azure.core.util.Context.NONE);
    }
}
