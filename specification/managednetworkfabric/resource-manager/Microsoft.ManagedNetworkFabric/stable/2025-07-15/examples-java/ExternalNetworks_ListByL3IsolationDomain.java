
/**
 * Samples for ExternalNetworks ListByL3IsolationDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/ExternalNetworks_ListByL3IsolationDomain.json
     */
    /**
     * Sample code: ExternalNetworks_ListByL3IsolationDomain_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksListByL3IsolationDomainMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.externalNetworks().listByL3IsolationDomain("example-rg", "example-externalnetwork",
            com.azure.core.util.Context.NONE);
    }
}
