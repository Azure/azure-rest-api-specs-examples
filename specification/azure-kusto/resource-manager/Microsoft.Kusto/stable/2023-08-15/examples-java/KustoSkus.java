/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSkus.json
     */
    /**
     * Sample code: KustoListRegionSkus.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoListRegionSkus(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.skus().list("westus", com.azure.core.util.Context.NONE);
    }
}
