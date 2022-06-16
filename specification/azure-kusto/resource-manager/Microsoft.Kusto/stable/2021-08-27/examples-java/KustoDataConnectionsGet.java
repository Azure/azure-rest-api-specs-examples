import com.azure.core.util.Context;

/** Samples for DataConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDataConnectionsGet.json
     */
    /**
     * Sample code: KustoDataConnectionsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .getWithResponse("kustorptest", "kustoclusterrptest4", "KustoDatabase8", "DataConnections8", Context.NONE);
    }
}
