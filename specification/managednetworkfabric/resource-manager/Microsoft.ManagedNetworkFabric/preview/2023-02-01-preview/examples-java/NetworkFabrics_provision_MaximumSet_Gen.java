/** Samples for NetworkFabrics Provision. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabrics_provision_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_provision_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsProvisionMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().provision("resourceGroupName", "FabricName", com.azure.core.util.Context.NONE);
    }
}
