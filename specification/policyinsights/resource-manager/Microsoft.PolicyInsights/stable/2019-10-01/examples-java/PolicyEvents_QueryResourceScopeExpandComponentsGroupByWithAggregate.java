
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/**
 * Samples for PolicyEvents ListQueryResultsForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyEvents_QueryResourceScopeExpandComponentsGroupByWithAggregate.json
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
        manager.policyEvents().listQueryResultsForResource(PolicyEventsResourceType.DEFAULT,
            "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName",
            null, null, null, null, null,
            "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
            null,
            "components($apply=groupby((tenantId, principalOid, policyDefinitionAction), aggregate($count as totalActions)))",
            null, com.azure.core.util.Context.NONE);
    }
}
