
/**
 * Samples for VirtualNetworkAddresses ListByCloudVmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * virtualNetworkAddresses_listByParent.json
     */
    /**
     * Sample code: List Virtual Network Addresses by VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listVirtualNetworkAddressesByVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.virtualNetworkAddresses().listByCloudVmCluster("rg000", "cluster1", com.azure.core.util.Context.NONE);
    }
}
