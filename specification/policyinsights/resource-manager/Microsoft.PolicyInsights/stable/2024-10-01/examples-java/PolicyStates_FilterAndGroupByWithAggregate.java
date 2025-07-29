
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;
import java.time.OffsetDateTime;

/**
 * Samples for PolicyStates ListQueryResultsForSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyStates_FilterAndGroupByWithAggregate.json
     */
    /**
     * Sample code: Filter and group with aggregate.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        filterAndGroupWithAggregate(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().listQueryResultsForSubscription(PolicyStatesResource.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", 2, "NumAuditDenyNonComplianceRecords desc", null,
            OffsetDateTime.parse("2019-10-05T18:00:00Z"), null,
            "IsCompliant eq false and (PolicyDefinitionAction eq 'audit' or PolicyDefinitionAction eq 'deny')",
            "groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId), aggregate($count as NumAuditDenyNonComplianceRecords))",
            null, com.azure.core.util.Context.NONE);
    }
}
