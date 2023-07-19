/** Samples for Clusters Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoClustersStart.json
     */
    /**
     * Sample code: KustoClustersStart.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersStart(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().start("kustorptest", "kustoCluster2", com.azure.core.util.Context.NONE);
    }
}
