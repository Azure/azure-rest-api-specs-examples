
/**
 * Samples for Clusters ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoClustersListSkus.json
     */
    /**
     * Sample code: KustoClustersListSkus.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersListSkus(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listSkus(com.azure.core.util.Context.NONE);
    }
}
