
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/**
 * Samples for PolicyEvents ListQueryResultsForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyEvents_QueryResourceScopeExpandComponents.json
     */
    /**
     * Sample code: Query components policy events for resource scope filtered by given assignment.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryComponentsPolicyEventsForResourceScopeFilteredByGivenAssignment(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyEvents().listQueryResultsForResource(PolicyEventsResourceType.DEFAULT,
            "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName",
            null, null, null, null, null,
            "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
            null, "components", null, com.azure.core.util.Context.NONE);
    }
}
