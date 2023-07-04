/** Samples for CloudServicesNetworks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/CloudServicesNetworks_ListByResourceGroup.json
     */
    /**
     * Sample code: List cloud services networks for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listCloudServicesNetworksForResourceGroup(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.cloudServicesNetworks().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
