/** Samples for KustoPoolPrincipalAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsGet.json
     */
    /**
     * Sample code: KustoPoolPrincipalAssignmentsGet.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolPrincipalAssignmentsGet(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolPrincipalAssignments()
            .getWithResponse(
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "kustoprincipal1",
                "kustorptest",
                com.azure.core.util.Context.NONE);
    }
}
