import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/** Samples for PolicyStates ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QuerySubscriptionScopeNextLink.json
     */
    /**
     * Sample code: Query latest at subscription scope with next link.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtSubscriptionScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForSubscription(
                PolicyStatesResource.LATEST,
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
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
