import com.azure.core.util.Context;

/** Samples for PolicyMetadata List. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyMetadata_List.json
     */
    /**
     * Sample code: Get collection of policy metadata resources.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void getCollectionOfPolicyMetadataResources(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyMetadatas().list(null, Context.NONE);
    }
}
