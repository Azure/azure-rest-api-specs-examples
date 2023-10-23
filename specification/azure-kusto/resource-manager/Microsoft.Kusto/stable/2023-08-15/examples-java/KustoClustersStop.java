/** Samples for Clusters Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClustersStop.json
     */
    /**
     * Sample code: KustoClustersStop.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersStop(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().stop("kustorptest", "kustoCluster2", com.azure.core.util.Context.NONE);
    }
}
