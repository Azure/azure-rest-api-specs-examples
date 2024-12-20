
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;
import java.time.OffsetDateTime;

/**
 * Samples for PolicyStates ListQueryResultsForSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyStates_FilterAndAggregateOnly.json
     */
    /**
     * Sample code: Filter and aggregate only.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void filterAndAggregateOnly(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().listQueryResultsForSubscription(PolicyStatesResource.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", null, null, null, OffsetDateTime.parse("2019-10-05T18:00:00Z"),
            null, "PolicyDefinitionAction eq 'deny'", "aggregate($count as NumDenyStates)", null,
            com.azure.core.util.Context.NONE);
    }
}
