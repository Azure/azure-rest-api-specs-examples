import com.azure.resourcemanager.kusto.models.Cluster;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoClustersUpdate.json
     */
    /**
     * Sample code: KustoClustersUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        Cluster resource =
            manager
                .clusters()
                .getByResourceGroupWithResponse("kustorptest", "kustoCluster2", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withIfMatch("*").apply();
    }
}
