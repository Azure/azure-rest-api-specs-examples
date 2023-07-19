/** Samples for DataConnections ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDataConnectionsListByDatabase.json
     */
    /**
     * Sample code: KustoDatabasesListByCluster.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesListByCluster(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .listByDatabase("kustorptest", "kustoCluster", "KustoDatabase8", com.azure.core.util.Context.NONE);
    }
}
