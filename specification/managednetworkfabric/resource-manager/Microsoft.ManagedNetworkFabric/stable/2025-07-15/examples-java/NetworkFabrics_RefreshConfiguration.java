
/**
 * Samples for NetworkFabrics RefreshConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_RefreshConfiguration.json
     */
    /**
     * Sample code: NetworkFabrics_RefreshConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsRefreshConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().refreshConfiguration("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}
