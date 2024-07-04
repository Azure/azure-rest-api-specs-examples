
/**
 * Samples for VirtualNetworkAddresses Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/virtualNetworkAddresses_get.json
     */
    /**
     * Sample code: Get Virtual Network Address.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        getVirtualNetworkAddress(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.virtualNetworkAddresses().getWithResponse("rg000", "cluster1", "hostname1",
            com.azure.core.util.Context.NONE);
    }
}
