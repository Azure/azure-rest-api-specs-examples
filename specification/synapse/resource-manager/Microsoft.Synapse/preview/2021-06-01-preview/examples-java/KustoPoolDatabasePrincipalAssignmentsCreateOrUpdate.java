import com.azure.resourcemanager.synapse.models.DatabasePrincipalRole;
import com.azure.resourcemanager.synapse.models.PrincipalType;

/** Samples for KustoPoolDatabasePrincipalAssignments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoPoolDatabasePrincipalAssignmentsCreateOrUpdate.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasePrincipalAssignmentsCreateOrUpdate(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabasePrincipalAssignments()
            .define("kustoprincipal1")
            .withExistingDatabase("synapseWorkspaceName", "kustoclusterrptest4", "Kustodatabase8", "kustorptest")
            .withPrincipalId("87654321-1234-1234-1234-123456789123")
            .withRole(DatabasePrincipalRole.ADMIN)
            .withTenantId("12345678-1234-1234-1234-123456789123")
            .withPrincipalType(PrincipalType.APP)
            .create();
    }
}
