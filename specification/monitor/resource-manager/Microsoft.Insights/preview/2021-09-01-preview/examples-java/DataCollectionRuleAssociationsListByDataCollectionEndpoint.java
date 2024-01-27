
import com.azure.core.util.Context;

/** Samples for DataCollectionRuleAssociations ListByDataCollectionEndpoint. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/
     * DataCollectionRuleAssociationsListByDataCollectionEndpoint.json
     */
    /**
     * Sample code: List associations for specified data collection endpoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAssociationsForSpecifiedDataCollectionEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getDataCollectionRuleAssociations()
            .listByDataCollectionEndpoint("myResourceGroup", "myDataCollectionEndpointName", Context.NONE);
    }
}
