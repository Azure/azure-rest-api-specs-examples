
/**
 * Samples for ExternalNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/ExternalNetworks_Delete.json
     */
    /**
     * Sample code: ExternalNetworks_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.externalNetworks().delete("example-rg", "example-externalnetwork", "example-ext",
            com.azure.core.util.Context.NONE);
    }
}
