import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyTrackedResourcesResourceType;

/** Samples for PolicyTrackedResources ListQueryResultsForResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceGroupScope.json
     */
    /**
     * Sample code: Query at resource group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtResourceGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyTrackedResources()
            .listQueryResultsForResourceGroup(
                "myResourceGroup", PolicyTrackedResourcesResourceType.DEFAULT, null, null, Context.NONE);
    }
}
