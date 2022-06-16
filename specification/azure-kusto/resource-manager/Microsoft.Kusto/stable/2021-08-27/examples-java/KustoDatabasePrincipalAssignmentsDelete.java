import com.azure.core.util.Context;

/** Samples for DatabasePrincipalAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabasePrincipalAssignmentsDelete.json
     */
    /**
     * Sample code: KustoDatabasePrincipalAssignmentsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasePrincipalAssignmentsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databasePrincipalAssignments()
            .delete("kustorptest", "kustoclusterrptest4", "Kustodatabase8", "kustoprincipal1", Context.NONE);
    }
}
