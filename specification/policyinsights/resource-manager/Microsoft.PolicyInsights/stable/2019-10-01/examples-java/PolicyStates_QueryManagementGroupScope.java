import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/** Samples for PolicyStates ListQueryResultsForManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryManagementGroupScope.json
     */
    /**
     * Sample code: Query latest at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForManagementGroup(
                PolicyStatesResource.LATEST,
                "myManagementGroup",
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
