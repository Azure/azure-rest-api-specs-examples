import com.azure.resourcemanager.kusto.models.ClusterCheckNameRequest;

/** Samples for Clusters CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoClustersCheckNameAvailability.json
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
                "westus", new ClusterCheckNameRequest().withName("kustoCluster"), com.azure.core.util.Context.NONE);
    }
}
