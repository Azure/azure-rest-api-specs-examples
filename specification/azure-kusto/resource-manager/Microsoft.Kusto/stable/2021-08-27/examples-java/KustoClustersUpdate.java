import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.Cluster;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClustersUpdate.json
     */
    /**
     * Sample code: KustoClustersUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        Cluster resource =
            manager
                .clusters()
                .getByResourceGroupWithResponse("kustorptest", "kustoclusterrptest4", Context.NONE)
                .getValue();
        resource.update().withIfMatch("*").apply();
    }
}
