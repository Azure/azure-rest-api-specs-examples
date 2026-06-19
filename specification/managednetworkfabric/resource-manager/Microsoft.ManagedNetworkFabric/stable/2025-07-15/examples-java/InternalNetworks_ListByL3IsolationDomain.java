
/**
 * Samples for InternalNetworks ListByL3IsolationDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternalNetworks_ListByL3IsolationDomain.json
     */
    /**
     * Sample code: InternalNetworks_ListByL3IsolationDomain_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksListByL3IsolationDomainMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internalNetworks().listByL3IsolationDomain("example-rg", "example-l3isd",
            com.azure.core.util.Context.NONE);
    }
}
