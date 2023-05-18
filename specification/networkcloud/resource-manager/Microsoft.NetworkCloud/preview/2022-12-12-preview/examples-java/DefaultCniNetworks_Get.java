/** Samples for DefaultCniNetworks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/DefaultCniNetworks_Get.json
     */
    /**
     * Sample code: Get default CNI network.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getDefaultCNINetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .defaultCniNetworks()
            .getByResourceGroupWithResponse(
                "resourceGroupName", "defaultCniNetworkName", com.azure.core.util.Context.NONE);
    }
}
