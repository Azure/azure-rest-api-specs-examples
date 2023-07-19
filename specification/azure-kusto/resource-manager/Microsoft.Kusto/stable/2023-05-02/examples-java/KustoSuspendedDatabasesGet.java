/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoSuspendedDatabasesGet.json
     */
    /**
     * Sample code: KustoSuspendedDatabasesGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoSuspendedDatabasesGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .getWithResponse("kustorptest", "kustoCluster", "KustoDatabase9", com.azure.core.util.Context.NONE);
    }
}
