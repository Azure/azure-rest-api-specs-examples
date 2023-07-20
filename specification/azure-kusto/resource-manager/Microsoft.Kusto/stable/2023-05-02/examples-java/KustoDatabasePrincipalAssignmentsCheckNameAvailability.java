import com.azure.resourcemanager.kusto.models.DatabasePrincipalAssignmentCheckNameRequest;

/** Samples for DatabasePrincipalAssignments CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoDatabasePrincipalAssignmentsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoDatabaseCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabaseCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databasePrincipalAssignments()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoCluster",
                "Kustodatabase8",
                new DatabasePrincipalAssignmentCheckNameRequest().withName("kustoprincipal1"),
                com.azure.core.util.Context.NONE);
    }
}
