import com.azure.core.util.Context;

/** Samples for PolicyEvents ListQueryResultsForPolicyDefinition. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QuerySubscriptionLevelPolicyDefinitionScopeNextLink.json
     */
    /**
     * Sample code: Query at subscription level policy definition scope with next link.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtSubscriptionLevelPolicyDefinitionScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForPolicyDefinition(
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                "24813039-7534-408a-9842-eb99f45721b1",
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                "WpmWfBSvPhkAK6QD",
                Context.NONE);
    }
}
