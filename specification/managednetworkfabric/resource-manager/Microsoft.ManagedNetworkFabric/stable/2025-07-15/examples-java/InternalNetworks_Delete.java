
/**
 * Samples for InternalNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternalNetworks_Delete.json
     */
    /**
     * Sample code: InternalNetworks_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internalNetworks().delete("example-rg", "example-l3isd", "example-internalnetwork",
            com.azure.core.util.Context.NONE);
    }
}
