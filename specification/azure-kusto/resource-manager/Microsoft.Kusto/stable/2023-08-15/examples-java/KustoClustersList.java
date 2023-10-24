/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClustersList.json
     */
    /**
     * Sample code: KustoClustersList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().list(com.azure.core.util.Context.NONE);
    }
}
