/** Samples for Databases ListPrincipals. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabaseListPrincipals.json
     */
    /**
     * Sample code: KustoDatabaseListPrincipals.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabaseListPrincipals(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databases()
            .listPrincipals("kustorptest", "kustoCluster", "KustoDatabase8", com.azure.core.util.Context.NONE);
    }
}
