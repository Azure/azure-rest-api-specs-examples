/** Samples for NetworkFabrics RefreshConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_refreshConfiguration_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_refreshConfiguration_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsRefreshConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().refreshConfiguration("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}
