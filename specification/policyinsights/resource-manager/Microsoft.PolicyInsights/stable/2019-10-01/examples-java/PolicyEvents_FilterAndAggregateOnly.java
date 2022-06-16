import com.azure.core.util.Context;
import java.time.OffsetDateTime;

/** Samples for PolicyEvents ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_FilterAndAggregateOnly.json
     */
    /**
     * Sample code: Filter and aggregate only.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void filterAndAggregateOnly(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForSubscription(
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                null,
                null,
                null,
                OffsetDateTime.parse("2018-02-05T18:00:00Z"),
                null,
                "PolicyDefinitionAction eq 'deny'",
                "aggregate($count as NumDenyEvents)",
                null,
                Context.NONE);
    }
}
