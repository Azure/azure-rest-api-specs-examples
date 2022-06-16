import com.azure.core.util.Context;

/** Samples for DataConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDataConnectionsDelete.json
     */
    /**
     * Sample code: KustoDataConnectionsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .delete("kustorptest", "kustoCluster", "KustoDatabase8", "dataConnectionTest", Context.NONE);
    }
}
