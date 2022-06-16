import com.azure.core.util.Context;

/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClustersList.json
     */
    /**
     * Sample code: KustoClustersList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().list(Context.NONE);
    }
}
