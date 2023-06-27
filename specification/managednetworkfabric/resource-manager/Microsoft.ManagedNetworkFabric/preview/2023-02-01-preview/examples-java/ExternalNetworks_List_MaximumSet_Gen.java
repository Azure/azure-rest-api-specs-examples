/** Samples for ExternalNetworks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/ExternalNetworks_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExternalNetworks_List_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksListMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.externalNetworks().list("resourceGroupName", "example-l3domain", com.azure.core.util.Context.NONE);
    }
}
