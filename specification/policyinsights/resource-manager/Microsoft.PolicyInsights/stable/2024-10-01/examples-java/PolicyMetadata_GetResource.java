
/**
 * Samples for PolicyMetadata GetResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/
     * PolicyMetadata_GetResource.json
     */
    /**
     * Sample code: Get a single policy metadata resource.
     * 
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        getASinglePolicyMetadataResource(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager.policyMetadatas().getResourceWithResponse("NIST_SP_800-53_R4_AC-2", com.azure.core.util.Context.NONE);
    }
}
