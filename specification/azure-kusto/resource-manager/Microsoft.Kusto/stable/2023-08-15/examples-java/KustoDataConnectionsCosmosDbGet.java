/** Samples for DataConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDataConnectionsCosmosDbGet.json
     */
    /**
     * Sample code: KustoDataConnectionsCosmosDbGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsCosmosDbGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .dataConnections()
            .getWithResponse(
                "kustorptest",
                "kustoCluster",
                "KustoDatabase1",
                "dataConnectionTest",
                com.azure.core.util.Context.NONE);
    }
}
