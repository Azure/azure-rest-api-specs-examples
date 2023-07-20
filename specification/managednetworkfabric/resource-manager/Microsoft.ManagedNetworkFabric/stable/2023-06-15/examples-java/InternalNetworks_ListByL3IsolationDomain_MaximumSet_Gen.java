/** Samples for InternalNetworks ListByL3IsolationDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternalNetworks_ListByL3IsolationDomain_MaximumSet_Gen.json
     */
    /**
     * Sample code: InternalNetworks_ListByL3IsolationDomain_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksListByL3IsolationDomainMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .internalNetworks()
            .listByL3IsolationDomain("example-rg", "example-l3domain", com.azure.core.util.Context.NONE);
    }
}
