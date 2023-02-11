/** Samples for KustoPoolDatabasePrincipalAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsGet.json
     */
    /**
     * Sample code: KustoPoolDatabasePrincipalAssignmentsGet.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasePrincipalAssignmentsGet(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabasePrincipalAssignments()
            .getWithResponse(
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "Kustodatabase8",
                "kustoprincipal1",
                "kustorptest",
                com.azure.core.util.Context.NONE);
    }
}
