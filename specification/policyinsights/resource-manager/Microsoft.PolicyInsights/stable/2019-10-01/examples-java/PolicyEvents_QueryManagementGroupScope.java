import com.azure.core.util.Context;

/** Samples for PolicyEvents ListQueryResultsForManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryManagementGroupScope.json
     */
    /**
     * Sample code: Query at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForManagementGroup(
                "myManagementGroup", null, null, null, null, null, null, null, null, Context.NONE);
    }
}
