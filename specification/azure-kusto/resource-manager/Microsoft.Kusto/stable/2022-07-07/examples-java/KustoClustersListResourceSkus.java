import com.azure.core.util.Context;

/** Samples for Clusters ListSkusByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoClustersListResourceSkus.json
     */
    /**
     * Sample code: KustoClustersListResourceSkus.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersListResourceSkus(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listSkusByResource("kustorptest", "kustoCluster", Context.NONE);
    }
}
