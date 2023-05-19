/** Samples for DefaultCniNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/DefaultCniNetworks_Delete.json
     */
    /**
     * Sample code: Delete default CNI network.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteDefaultCNINetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .defaultCniNetworks()
            .delete("resourceGroupName", "defaultCniNetworkName", com.azure.core.util.Context.NONE);
    }
}
