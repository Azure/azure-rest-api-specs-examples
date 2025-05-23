
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;

/**
 * Samples for PolicyEvents ListQueryResultsForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyEvents_QueryResourceScopeNextLink.json
     */
    /**
     * Sample code: Query at resource scope with next link.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        queryAtResourceScopeWithNextLink(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyEvents().listQueryResultsForResource(PolicyEventsResourceType.DEFAULT,
            "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName",
            null, null, null, null, null, null, null, null, "WpmWfBSvPhkAK6QD", com.azure.core.util.Context.NONE);
    }
}
