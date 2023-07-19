/** Samples for NetworkFabrics Provision. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_provision_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_provision_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsProvisionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().provision("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}
