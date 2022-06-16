import com.azure.core.util.Context;
import java.time.OffsetDateTime;

/** Samples for PolicyStates SummarizeForPolicySetDefinition. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicySetDefinitionScope.json
     */
    /**
     * Sample code: Summarize at policy set definition scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void summarizeAtPolicySetDefinitionScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .summarizeForPolicySetDefinitionWithResponse(
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                "3e3807c1-65c9-49e0-a406-82d8ae3e338c",
                1,
                OffsetDateTime.parse("2019-10-05T18:00:00Z"),
                OffsetDateTime.parse("2019-10-06T18:00:00Z"),
                "PolicyDefinitionAction eq 'deny'",
                Context.NONE);
    }
}
