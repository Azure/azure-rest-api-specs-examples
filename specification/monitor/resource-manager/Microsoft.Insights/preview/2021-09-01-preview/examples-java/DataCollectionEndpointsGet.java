
import com.azure.core.util.Context;

/** Samples for DataCollectionEndpoints GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionEndpointsGet.json
     */
    /**
     * Sample code: Get data collection endpoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDataCollectionEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionEndpoints()
            .getByResourceGroupWithResponse("myResourceGroup", "myCollectionEndpoint", Context.NONE);
    }
}
