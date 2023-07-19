/** Samples for ClusterPrincipalAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoClusterPrincipalAssignmentsDelete.json
     */
    /**
     * Sample code: KustoClusterPrincipalAssignmentsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterPrincipalAssignmentsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusterPrincipalAssignments()
            .delete("kustorptest", "kustoCluster", "kustoprincipal1", com.azure.core.util.Context.NONE);
    }
}
