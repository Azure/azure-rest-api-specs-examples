import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.DatabasePrincipalAssignmentCheckNameRequest;

/** Samples for KustoPoolDatabasePrincipalAssignments CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoPoolDatabasePrincipalAssignmentsCheckNameAvailability.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasePrincipalAssignmentsCheckNameAvailability(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabasePrincipalAssignments()
            .checkNameAvailabilityWithResponse(
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "Kustodatabase8",
                "kustorptest",
                new DatabasePrincipalAssignmentCheckNameRequest().withName("kustoprincipal1"),
                Context.NONE);
    }
}
