import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.ClusterPrincipalAssignmentCheckNameRequest;

/** Samples for KustoPoolPrincipalAssignments CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrincipalAssignmentsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoPoolPrincipalAssignmentsCheckNameAvailability.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolPrincipalAssignmentsCheckNameAvailability(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolPrincipalAssignments()
            .checkNameAvailabilityWithResponse(
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "kustorptest",
                new ClusterPrincipalAssignmentCheckNameRequest().withName("kustoprincipal1"),
                Context.NONE);
    }
}
