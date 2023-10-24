import com.azure.resourcemanager.kusto.models.ClusterPrincipalRole;
import com.azure.resourcemanager.kusto.models.PrincipalType;

/** Samples for ClusterPrincipalAssignments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClusterPrincipalAssignmentsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoClusterPrincipalAssignmentsCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterPrincipalAssignmentsCreateOrUpdate(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusterPrincipalAssignments()
            .define("kustoprincipal1")
            .withExistingCluster("kustorptest", "kustoCluster")
            .withPrincipalId("87654321-1234-1234-1234-123456789123")
            .withRole(ClusterPrincipalRole.ALL_DATABASES_ADMIN)
            .withTenantId("12345678-1234-1234-1234-123456789123")
            .withPrincipalType(PrincipalType.APP)
            .create();
    }
}
