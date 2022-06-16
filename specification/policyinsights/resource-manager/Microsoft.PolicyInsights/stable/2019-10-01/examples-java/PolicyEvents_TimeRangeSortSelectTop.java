import com.azure.core.util.Context;
import java.time.OffsetDateTime;

/** Samples for PolicyEvents ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_TimeRangeSortSelectTop.json
     */
    /**
     * Sample code: Time range; sort, select and limit.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void timeRangeSortSelectAndLimit(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForSubscription(
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                2,
                "Timestamp desc, PolicyAssignmentId asc, SubscriptionId asc, ResourceGroup asc, ResourceId",
                "Timestamp, PolicyAssignmentId, PolicyDefinitionId, SubscriptionId, ResourceGroup, ResourceId",
                OffsetDateTime.parse("2018-02-05T18:00:00Z"),
                OffsetDateTime.parse("2018-02-06T18:00:00Z"),
                null,
                null,
                null,
                Context.NONE);
    }
}
