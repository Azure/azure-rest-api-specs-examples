import com.azure.core.util.Context;

/** Samples for ClusterPrincipalAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClusterPrincipalAssignmentsDelete.json
     */
    /**
     * Sample code: KustoClusterPrincipalAssignmentsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterPrincipalAssignmentsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusterPrincipalAssignments()
            .delete("kustorptest", "kustoclusterrptest4", "kustoprincipal1", Context.NONE);
    }
}
