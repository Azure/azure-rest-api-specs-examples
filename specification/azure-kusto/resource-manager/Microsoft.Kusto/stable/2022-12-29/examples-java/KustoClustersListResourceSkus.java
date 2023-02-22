/** Samples for Clusters ListSkusByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoClustersListResourceSkus.json
     */
    /**
     * Sample code: KustoClustersListResourceSkus.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersListResourceSkus(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listSkusByResource("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
