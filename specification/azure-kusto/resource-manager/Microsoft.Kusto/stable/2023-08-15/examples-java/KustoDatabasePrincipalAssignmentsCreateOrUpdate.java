import com.azure.resourcemanager.kusto.models.DatabasePrincipalRole;
import com.azure.resourcemanager.kusto.models.PrincipalType;

/** Samples for DatabasePrincipalAssignments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabasePrincipalAssignmentsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoDatabasePrincipalAssignmentsCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasePrincipalAssignmentsCreateOrUpdate(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .databasePrincipalAssignments()
            .define("kustoprincipal1")
            .withExistingDatabase("kustorptest", "kustoCluster", "Kustodatabase8")
            .withPrincipalId("87654321-1234-1234-1234-123456789123")
            .withRole(DatabasePrincipalRole.ADMIN)
            .withTenantId("12345678-1234-1234-1234-123456789123")
            .withPrincipalType(PrincipalType.APP)
            .create();
    }
}
