/** Samples for VirtualNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_Delete.json
     */
    /**
     * Sample code: VirtualNetworks_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void virtualNetworksDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .virtualNetworks()
            .delete("resourceGroupName", "{labName}", "{virtualNetworkName}", com.azure.core.util.Context.NONE);
    }
}
