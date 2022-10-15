import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;
import java.time.OffsetDateTime;

/** Samples for PolicyEvents ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_FilterAndGroupByWithoutAggregate.json
     */
    /**
     * Sample code: Filter and group without aggregate.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void filterAndGroupWithoutAggregate(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForSubscription(
                PolicyEventsResourceType.DEFAULT,
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                2,
                null,
                null,
                OffsetDateTime.parse("2018-01-05T18:00:00Z"),
                null,
                "PolicyDefinitionAction ne 'audit' and PolicyDefinitionAction ne 'append'",
                "groupby((PolicyAssignmentId, PolicyDefinitionId, PolicyDefinitionAction, ResourceId))",
                null,
                Context.NONE);
    }
}
