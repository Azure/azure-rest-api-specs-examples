
/**
 * Samples for ExternalNetworks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/ExternalNetworks_Get.json
     */
    /**
     * Sample code: ExternalNetworks_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void externalNetworksGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.externalNetworks().getWithResponse("example-rg", "example-externalnetwork", "example-ext",
            com.azure.core.util.Context.NONE);
    }
}
