import com.azure.core.util.Context;

/** Samples for DataConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDataConnectionsEventGridGet.json
     */
    /**
     * Sample code: KustoDataConnectionsEventGridGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsEventGridGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .getWithResponse("kustorptest", "kustoCluster", "KustoDatabase8", "dataConnectionTest", Context.NONE);
    }
}
