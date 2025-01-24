
/**
 * Samples for Databases ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDatabasesListByCluster
     * .json
     */
    /**
     * Sample code: KustoDatabasesListByCluster.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesListByCluster(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.databases().listByCluster("kustorptest", "kustoCluster", null, null, com.azure.core.util.Context.NONE);
    }
}
