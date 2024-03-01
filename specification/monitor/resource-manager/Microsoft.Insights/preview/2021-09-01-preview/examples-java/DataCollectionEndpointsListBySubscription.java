
/**
 * Samples for DataCollectionEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionEndpointsListBySubscription.json
     */
    /**
     * Sample code: List data collection endpoints by subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDataCollectionEndpointsBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionEndpoints()
            .list(com.azure.core.util.Context.NONE);
    }
}
