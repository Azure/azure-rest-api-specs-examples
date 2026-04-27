
/**
 * Samples for Organization DeleteClusterApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_DeleteClusterAPIKey_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organization_DeleteClusterAPIKey_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationDeleteClusterAPIKeyMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().deleteClusterApiKeyWithResponse("rgconfluent", "lokrfxecjwbnejqluwbwqcairu",
            "lqyopqadqide", com.azure.core.util.Context.NONE);
    }
}
