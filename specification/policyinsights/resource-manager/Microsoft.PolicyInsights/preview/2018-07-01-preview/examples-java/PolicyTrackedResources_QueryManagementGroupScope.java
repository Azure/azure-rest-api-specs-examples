import com.azure.core.util.Context;

/** Samples for PolicyTrackedResources ListQueryResultsForManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryManagementGroupScope.json
     */
    /**
     * Sample code: Query at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyTrackedResources()
            .listQueryResultsForManagementGroup("myManagementGroup", null, null, Context.NONE);
    }
}
