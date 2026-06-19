
/**
 * Samples for NetworkFabrics Deprovision.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_Deprovision.json
     */
    /**
     * Sample code: NetworkFabrics_Deprovision_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsDeprovisionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().deprovision("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}
