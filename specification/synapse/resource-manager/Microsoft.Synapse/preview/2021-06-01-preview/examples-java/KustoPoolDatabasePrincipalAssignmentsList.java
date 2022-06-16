import com.azure.core.util.Context;

/** Samples for KustoPoolDatabasePrincipalAssignments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsList.json
     */
    /**
     * Sample code: KustoPoolDatabasePrincipalAssignmentsList.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasePrincipalAssignmentsList(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabasePrincipalAssignments()
            .list("synapseWorkspaceName", "kustoclusterrptest4", "Kustodatabase8", "kustorptest", Context.NONE);
    }
}
