
/**
 * Samples for DataCollectionEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionEndpointsDelete.json
     */
    /**
     * Sample code: Delete data collection endpoint.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteDataCollectionEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionEndpoints()
            .deleteWithResponse("myResourceGroup", "myCollectionEndpoint", com.azure.core.util.Context.NONE);
    }
}
