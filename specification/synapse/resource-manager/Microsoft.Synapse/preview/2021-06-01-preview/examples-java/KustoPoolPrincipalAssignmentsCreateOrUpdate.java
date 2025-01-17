
import com.azure.resourcemanager.synapse.models.ClusterPrincipalRole;
import com.azure.resourcemanager.synapse.models.PrincipalType;

/**
 * Samples for KustoPoolPrincipalAssignments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/
     * KustoPoolPrincipalAssignmentsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoPoolPrincipalAssignmentsCreateOrUpdate.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        kustoPoolPrincipalAssignmentsCreateOrUpdate(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.kustoPoolPrincipalAssignments().define("kustoprincipal1")
            .withExistingKustoPool("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest")
            .withPrincipalId("87654321-1234-1234-1234-123456789123").withRole(ClusterPrincipalRole.ALL_DATABASES_ADMIN)
            .withTenantId("12345678-1234-1234-1234-123456789123").withPrincipalType(PrincipalType.APP).create();
    }
}
