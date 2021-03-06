import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.ClusterPrincipalAssignmentCheckNameRequest;

/** Samples for ClusterPrincipalAssignments CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClusterPrincipalAssignmentsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoClusterPrincipalAssignmentsCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterPrincipalAssignmentsCheckNameAvailability(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusterPrincipalAssignments()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoclusterrptest4",
                new ClusterPrincipalAssignmentCheckNameRequest().withName("kustoprincipal1"),
                Context.NONE);
    }
}
