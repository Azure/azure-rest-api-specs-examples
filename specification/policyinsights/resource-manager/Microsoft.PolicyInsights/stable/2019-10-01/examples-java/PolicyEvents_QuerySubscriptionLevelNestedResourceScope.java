import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/** Samples for PolicyEvents ListQueryResultsForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QuerySubscriptionLevelNestedResourceScope.json
     */
    /**
     * Sample code: Query at subscription level nested resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtSubscriptionLevelNestedResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForResource(
                PolicyEventsResourceType.DEFAULT,
                "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.SomeNamespace/someResourceType/someResource/someNestedResourceType/someNestedResource",
                null,
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

    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceScopeExpandComponentsGroupByWithAggregate.json
     */
    /**
     * Sample code: Query components policy events count grouped by user and action type for resource scope filtered by
     * given assignment.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        queryComponentsPolicyEventsCountGroupedByUserAndActionTypeForResourceScopeFilteredByGivenAssignment(
            com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForResource(
                PolicyEventsResourceType.DEFAULT,
                "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName",
                null,
                null,
                null,
                null,
                null,
                "policyAssignmentId eq"
                    + " '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
                null,
                "components($apply=groupby((tenantId, principalOid, policyDefinitionAction), aggregate($count as"
                    + " totalActions)))",
                null,
                Context.NONE);
    }

    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceScopeNextLink.json
     */
    /**
     * Sample code: Query at resource scope with next link.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtResourceScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForResource(
                PolicyEventsResourceType.DEFAULT,
                "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName",
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                "WpmWfBSvPhkAK6QD",
                Context.NONE);
    }
}
