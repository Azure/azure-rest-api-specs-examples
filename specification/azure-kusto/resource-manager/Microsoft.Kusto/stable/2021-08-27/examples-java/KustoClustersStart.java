import com.azure.core.util.Context;

/** Samples for Clusters Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClustersStart.json
     */
    /**
     * Sample code: KustoClustersStart.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersStart(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().start("kustorptest", "kustoclusterrptest4", Context.NONE);
    }
}
