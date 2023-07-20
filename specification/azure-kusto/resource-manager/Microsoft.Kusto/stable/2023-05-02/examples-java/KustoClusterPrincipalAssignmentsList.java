/** Samples for ClusterPrincipalAssignments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoClusterPrincipalAssignmentsList.json
     */
    /**
     * Sample code: KustoPrincipalAssignmentsList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoPrincipalAssignmentsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusterPrincipalAssignments().list("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
