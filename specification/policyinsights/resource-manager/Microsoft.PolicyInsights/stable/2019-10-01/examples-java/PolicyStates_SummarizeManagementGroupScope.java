import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyStatesSummaryResourceType;
import java.time.OffsetDateTime;

/** Samples for PolicyStates SummarizeForManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeManagementGroupScope.json
     */
    /**
     * Sample code: Summarize at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void summarizeAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .summarizeForManagementGroupWithResponse(
                PolicyStatesSummaryResourceType.LATEST,
                "myManagementGroup",
                0,
                OffsetDateTime.parse("2019-10-05T18:00:00Z"),
                OffsetDateTime.parse("2019-10-06T18:00:00Z"),
                "PolicyDefinitionAction eq 'deny' or PolicyDefinitionAction eq 'audit'",
                Context.NONE);
    }
}
