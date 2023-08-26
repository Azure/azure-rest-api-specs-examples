/** Samples for TrunkedNetworks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/TrunkedNetworks_ListByResourceGroup.json
     */
    /**
     * Sample code: List Trunked networks for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listTrunkedNetworksForResourceGroup(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.trunkedNetworks().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
