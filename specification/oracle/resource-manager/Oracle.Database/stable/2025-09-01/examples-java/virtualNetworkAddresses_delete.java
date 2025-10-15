
/**
 * Samples for VirtualNetworkAddresses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/virtualNetworkAddresses_delete.json
     */
    /**
     * Sample code: VirtualNetworkAddresses_delete.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        virtualNetworkAddressesDelete(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.virtualNetworkAddresses().delete("rg000", "cluster1", "hostname1", com.azure.core.util.Context.NONE);
    }
}
