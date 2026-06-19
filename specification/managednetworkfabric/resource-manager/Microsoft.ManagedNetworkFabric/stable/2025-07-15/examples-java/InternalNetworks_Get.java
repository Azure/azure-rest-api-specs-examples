
/**
 * Samples for InternalNetworks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/InternalNetworks_Get.json
     */
    /**
     * Sample code: InternalNetworks_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void internalNetworksGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.internalNetworks().getWithResponse("example-rg", "example-l3isd", "example-internalnetwork",
            com.azure.core.util.Context.NONE);
    }
}
