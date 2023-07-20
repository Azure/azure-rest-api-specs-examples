/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDatabasesGet.json
     */
    /**
     * Sample code: KustoDatabasesGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .getWithResponse("kustorptest", "kustoCluster", "KustoDatabase8", com.azure.core.util.Context.NONE);
    }
}
