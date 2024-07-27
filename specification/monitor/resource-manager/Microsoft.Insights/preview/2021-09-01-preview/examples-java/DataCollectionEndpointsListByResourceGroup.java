
/**
 * Samples for DataCollectionEndpoints ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionEndpointsListByResourceGroup.json
     */
    /**
     * Sample code: List data collection endpoints by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listDataCollectionEndpointsByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionEndpoints()
            .listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
