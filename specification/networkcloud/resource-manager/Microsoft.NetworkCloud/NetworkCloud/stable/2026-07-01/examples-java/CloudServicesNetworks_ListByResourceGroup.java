
/**
 * Samples for CloudServicesNetworks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/CloudServicesNetworks_ListByResourceGroup.json
     */
    /**
     * Sample code: List cloud services networks for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listCloudServicesNetworksForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.cloudServicesNetworks().listByResourceGroup("resourceGroupName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
