/** Samples for Databases Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasesDelete.json
     */
    /**
     * Sample code: KustoDatabasesDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.databases().delete("kustorptest", "kustoCluster", "KustoDatabase8", com.azure.core.util.Context.NONE);
    }
}
