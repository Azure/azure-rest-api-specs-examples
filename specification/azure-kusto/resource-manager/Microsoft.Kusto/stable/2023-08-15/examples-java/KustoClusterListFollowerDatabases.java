
/**
 * Samples for Clusters ListFollowerDatabases.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/
     * KustoClusterListFollowerDatabases.json
     */
    /**
     * Sample code: KustoClusterListFollowerDatabases.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterListFollowerDatabases(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listFollowerDatabases("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
