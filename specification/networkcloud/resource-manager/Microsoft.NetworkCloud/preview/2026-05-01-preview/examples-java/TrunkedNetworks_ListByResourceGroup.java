
/**
 * Samples for TrunkedNetworks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/TrunkedNetworks_ListByResourceGroup.json
     */
    /**
     * Sample code: List Trunked networks for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listTrunkedNetworksForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.trunkedNetworks().listByResourceGroup("resourceGroupName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
