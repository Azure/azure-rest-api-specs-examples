/** Samples for NetworkFabrics Deprovision. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabrics_deprovision_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_deprovision_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsDeprovisionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().deprovision("resourceGroupName", "FabricName", com.azure.core.util.Context.NONE);
    }
}
