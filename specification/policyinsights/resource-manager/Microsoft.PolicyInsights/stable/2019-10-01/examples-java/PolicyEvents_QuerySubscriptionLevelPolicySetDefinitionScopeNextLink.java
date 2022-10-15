import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/** Samples for PolicyEvents ListQueryResultsForPolicySetDefinition. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QuerySubscriptionLevelPolicySetDefinitionScopeNextLink.json
     */
    /**
     * Sample code: Query at subscription level policy set definition scope with next link.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAtSubscriptionLevelPolicySetDefinitionScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForPolicySetDefinition(
                PolicyEventsResourceType.DEFAULT,
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                "3e3807c1-65c9-49e0-a406-82d8ae3e338c",
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
