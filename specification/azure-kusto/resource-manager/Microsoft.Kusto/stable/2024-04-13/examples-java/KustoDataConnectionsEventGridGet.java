
/**
 * Samples for DataConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoDataConnectionsEventGridGet.json
     */
    /**
     * Sample code: KustoDataConnectionsEventGridGet.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDataConnectionsEventGridGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.dataConnections().getWithResponse("kustorptest", "kustoCluster", "KustoDatabase8", "dataConnectionTest",
            com.azure.core.util.Context.NONE);
    }
}
