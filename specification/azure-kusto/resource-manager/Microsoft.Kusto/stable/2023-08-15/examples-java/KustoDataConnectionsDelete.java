/** Samples for DataConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionsDelete.json
     */
    /**
     * Sample code: KustoDataConnectionsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .delete(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase8",
                "dataConnectionTest",
                com.azure.core.util.Context.NONE);
    }
}
