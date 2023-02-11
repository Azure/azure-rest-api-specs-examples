/** Samples for KustoPoolPrincipalAssignments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsList.json
     */
    /**
     * Sample code: KustoPoolPrincipalAssignmentsList.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolPrincipalAssignmentsList(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolPrincipalAssignments()
            .list("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", com.azure.core.util.Context.NONE);
    }
}
