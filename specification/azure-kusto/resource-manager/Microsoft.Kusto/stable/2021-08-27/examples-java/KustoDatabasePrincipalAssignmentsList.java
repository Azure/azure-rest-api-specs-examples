import com.azure.core.util.Context;

/** Samples for DatabasePrincipalAssignments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabasePrincipalAssignmentsList.json
     */
    /**
     * Sample code: KustoPrincipalAssignmentsList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoPrincipalAssignmentsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databasePrincipalAssignments()
            .list("kustorptest", "kustoclusterrptest4", "Kustodatabase8", Context.NONE);
    }
}
