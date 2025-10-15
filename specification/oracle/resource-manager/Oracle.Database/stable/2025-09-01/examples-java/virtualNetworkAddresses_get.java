
/**
 * Samples for VirtualNetworkAddresses Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/virtualNetworkAddresses_get.json
     */
    /**
     * Sample code: VirtualNetworkAddresses_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        virtualNetworkAddressesGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.virtualNetworkAddresses().getWithResponse("rg000", "cluster1", "hostname1",
            com.azure.core.util.Context.NONE);
    }
}
