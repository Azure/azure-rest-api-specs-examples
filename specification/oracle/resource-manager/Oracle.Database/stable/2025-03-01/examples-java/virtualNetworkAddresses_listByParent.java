
/**
 * Samples for VirtualNetworkAddresses ListByCloudVmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/virtualNetworkAddresses_listByParent.json
     */
    /**
     * Sample code: VirtualNetworkAddresses_listByCloudVmCluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void virtualNetworkAddressesListByCloudVmCluster(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.virtualNetworkAddresses().listByCloudVmCluster("rg000", "cluster1", com.azure.core.util.Context.NONE);
    }
}
