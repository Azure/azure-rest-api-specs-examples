import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/** Samples for PolicyStates ListQueryResultsForPolicyDefinition. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QuerySubscriptionLevelPolicyDefinitionScopeNextLink.json
     */
    /**
     * Sample code: Query latest at subscription level policy definition scope with next link.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtSubscriptionLevelPolicyDefinitionScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForPolicyDefinition(
                PolicyStatesResource.LATEST,
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
