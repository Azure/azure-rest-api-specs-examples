import com.azure.core.util.Context;

/** Samples for Clusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClustersListByResourceGroup.json
     */
    /**
     * Sample code: KustoClustersListByResourceGroup.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersListByResourceGroup(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listByResourceGroup("kustorptest", Context.NONE);
    }
}
