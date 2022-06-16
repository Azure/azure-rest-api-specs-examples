import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.ClusterCheckNameRequest;

/** Samples for Clusters CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClustersCheckNameAvailability.json
     */
    /**
     * Sample code: KustoClustersCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClustersCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusters()
            .checkNameAvailabilityWithResponse(
                "westus", new ClusterCheckNameRequest().withName("kustoclusterrptest4"), Context.NONE);
    }
}
