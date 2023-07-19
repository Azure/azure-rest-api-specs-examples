/** Samples for NetworkFabrics Deprovision. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_deprovision_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_deprovision_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsDeprovisionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().deprovision("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}
