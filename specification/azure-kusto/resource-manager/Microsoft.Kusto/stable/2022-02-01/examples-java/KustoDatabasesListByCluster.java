import com.azure.core.util.Context;

/** Samples for Databases ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoDatabasesListByCluster.json
     */
    /**
     * Sample code: KustoDatabasesListByCluster.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesListByCluster(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.databases().listByCluster("kustorptest", "kustoCluster", Context.NONE);
    }
}
