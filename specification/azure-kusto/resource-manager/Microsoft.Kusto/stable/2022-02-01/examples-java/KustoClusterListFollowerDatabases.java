import com.azure.core.util.Context;

/** Samples for Clusters ListFollowerDatabases. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClusterListFollowerDatabases.json
     */
    /**
     * Sample code: KustoClusterListFollowerDatabases.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterListFollowerDatabases(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listFollowerDatabases("kustorptest", "kustoCluster", Context.NONE);
    }
}
