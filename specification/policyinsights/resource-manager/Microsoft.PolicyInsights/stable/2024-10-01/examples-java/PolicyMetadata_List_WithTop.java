
/**
 * Samples for PolicyMetadata List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyMetadata_List_WithTop.json
     */
    /**
     * Sample code: Get collection of policy metadata resources using top query parameter.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void getCollectionOfPolicyMetadataResourcesUsingTopQueryParameter(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyMetadatas().list(1, com.azure.core.util.Context.NONE);
    }
}
