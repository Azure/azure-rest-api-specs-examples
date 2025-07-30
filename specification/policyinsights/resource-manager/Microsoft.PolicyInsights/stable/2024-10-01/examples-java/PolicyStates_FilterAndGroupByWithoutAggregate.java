
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;
import java.time.OffsetDateTime;

/**
 * Samples for PolicyStates ListQueryResultsForSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyStates_FilterAndGroupByWithoutAggregate.json
     */
    /**
     * Sample code: Filter and group without aggregate.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        filterAndGroupWithoutAggregate(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().listQueryResultsForSubscription(PolicyStatesResource.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", 2, null, null, OffsetDateTime.parse("2019-10-05T18:00:00Z"), null,
            "IsCompliant eq false and (PolicyDefinitionAction ne 'audit' and PolicyDefinitionAction ne 'append')",
            "groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId))", null,
            com.azure.core.util.Context.NONE);
    }
}
