
import com.azure.resourcemanager.policyinsights.models.PolicyTrackedResourcesResourceType;

/**
 * Samples for PolicyTrackedResources ListQueryResultsForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/
     * PolicyTrackedResources_QueryResourceScope.json
     */
    /**
     * Sample code: Query at resource scope.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtResourceScope(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyTrackedResources().listQueryResultsForResource(
            "subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource",
            PolicyTrackedResourcesResourceType.DEFAULT, null, null, com.azure.core.util.Context.NONE);
    }
}
