import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;
import java.time.OffsetDateTime;

/** Samples for PolicyStates ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_TimeRangeSortSelectTop.json
     */
    /**
     * Sample code: Time range; sort, select and limit.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void timeRangeSortSelectAndLimit(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForSubscription(
                PolicyStatesResource.LATEST,
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                2,
                "Timestamp desc, PolicyAssignmentId asc, SubscriptionId asc, ResourceGroup asc, ResourceId",
                "Timestamp, PolicyAssignmentId, PolicyDefinitionId, SubscriptionId, ResourceGroup, ResourceId,"
                    + " policyDefinitionGroupNames",
                OffsetDateTime.parse("2019-10-05T18:00:00Z"),
                OffsetDateTime.parse("2019-10-06T18:00:00Z"),
                null,
                null,
                null,
                Context.NONE);
    }
}
