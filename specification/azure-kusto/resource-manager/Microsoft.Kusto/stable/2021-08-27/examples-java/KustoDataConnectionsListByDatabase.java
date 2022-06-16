import com.azure.core.util.Context;

/** Samples for DataConnections ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDataConnectionsListByDatabase.json
     */
    /**
     * Sample code: KustoDatabasesListByCluster.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesListByCluster(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.dataConnections().listByDatabase("kustorptest", "kustoclusterrptest4", "KustoDatabase8", Context.NONE);
    }
}
