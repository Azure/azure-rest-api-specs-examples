import com.azure.core.util.Context;

/** Samples for DatabasePrincipalAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabasePrincipalAssignmentsGet.json
     */
    /**
     * Sample code: KustoDatabasePrincipalAssignmentsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasePrincipalAssignmentsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databasePrincipalAssignments()
            .getWithResponse("kustorptest", "kustoclusterrptest4", "Kustodatabase8", "kustoprincipal1", Context.NONE);
    }
}
