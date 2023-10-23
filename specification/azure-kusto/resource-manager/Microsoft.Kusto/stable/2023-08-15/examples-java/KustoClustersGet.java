/** Samples for Clusters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClustersGet.json
     */
    /**
     * Sample code: KustoClustersGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusters()
            .getByResourceGroupWithResponse("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
