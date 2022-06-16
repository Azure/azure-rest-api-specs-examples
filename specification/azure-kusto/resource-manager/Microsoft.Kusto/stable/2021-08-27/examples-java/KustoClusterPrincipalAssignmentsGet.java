import com.azure.core.util.Context;

/** Samples for ClusterPrincipalAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClusterPrincipalAssignmentsGet.json
     */
    /**
     * Sample code: KustoClusterPrincipalAssignmentsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterPrincipalAssignmentsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusterPrincipalAssignments()
            .getWithResponse("kustorptest", "kustoclusterrptest4", "kustoprincipal1", Context.NONE);
    }
}
