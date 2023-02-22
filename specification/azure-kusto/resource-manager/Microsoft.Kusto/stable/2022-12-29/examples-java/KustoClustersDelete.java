/** Samples for Clusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoClustersDelete.json
     */
    /**
     * Sample code: KustoClustersDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().delete("kustorptest", "kustoCluster2", com.azure.core.util.Context.NONE);
    }
}
