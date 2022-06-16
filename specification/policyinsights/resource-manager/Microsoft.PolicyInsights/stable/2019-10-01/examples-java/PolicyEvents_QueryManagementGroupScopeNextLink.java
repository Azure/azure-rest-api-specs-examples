import com.azure.core.util.Context;

/** Samples for PolicyEvents ListQueryResultsForManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryManagementGroupScopeNextLink.json
     */
    /**
     * Sample code: Query at management group scope with next link.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtManagementGroupScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForManagementGroup(
                "myManagementGroup", null, null, null, null, null, null, null, "WpmWfBSvPhkAK6QD", Context.NONE);
    }
}
