
/**
 * Samples for TrunkedNetworks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * TrunkedNetworks_ListByResourceGroup.json
     */
    /**
     * Sample code: List Trunked networks for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listTrunkedNetworksForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.trunkedNetworks().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
