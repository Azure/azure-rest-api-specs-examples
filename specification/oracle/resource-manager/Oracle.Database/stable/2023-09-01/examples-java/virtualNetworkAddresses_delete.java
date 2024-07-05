
/**
 * Samples for VirtualNetworkAddresses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/virtualNetworkAddresses_delete.
     * json
     */
    /**
     * Sample code: Delete Virtual Network Address.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        deleteVirtualNetworkAddress(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.virtualNetworkAddresses().delete("rg000", "cluster1", "hostname1", com.azure.core.util.Context.NONE);
    }
}
