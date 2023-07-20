/** Samples for DatabasePrincipalAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDatabasePrincipalAssignmentsGet.json
     */
    /**
     * Sample code: KustoDatabasePrincipalAssignmentsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasePrincipalAssignmentsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databasePrincipalAssignments()
            .getWithResponse(
                "kustorptest", "kustoCluster", "Kustodatabase8", "kustoprincipal1", com.azure.core.util.Context.NONE);
    }
}
