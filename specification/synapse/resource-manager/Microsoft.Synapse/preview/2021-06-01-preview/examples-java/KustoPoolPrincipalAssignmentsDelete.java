import com.azure.core.util.Context;

/** Samples for KustoPoolPrincipalAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsDelete.json
     */
    /**
     * Sample code: KustoPoolPrincipalAssignmentsDelete.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolPrincipalAssignmentsDelete(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolPrincipalAssignments()
            .delete("synapseWorkspaceName", "kustoclusterrptest4", "kustoprincipal1", "kustorptest", Context.NONE);
    }
}
