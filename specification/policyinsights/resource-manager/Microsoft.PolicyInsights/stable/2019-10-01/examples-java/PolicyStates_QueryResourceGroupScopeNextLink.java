
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/**
 * Samples for PolicyStates ListQueryResultsForResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/
     * PolicyStates_QueryResourceGroupScopeNextLink.json
     */
    /**
     * Sample code: Query latest at resource group scope with next link.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryLatestAtResourceGroupScopeWithNextLink(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyStates().listQueryResultsForResourceGroup(PolicyStatesResource.LATEST,
            "fffedd8f-ffff-fffd-fffd-fffed2f84852", "myResourceGroup", null, null, null, null, null, null, null,
            "WpmWfBSvPhkAK6QD", com.azure.core.util.Context.NONE);
    }
}
