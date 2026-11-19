
/**
 * Samples for Organization DeleteClusterApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/Organization_DeleteClusterAPIKey_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_DeleteClusterAPIKey_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationDeleteClusterAPIKeyMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().deleteClusterApiKeyWithResponse("rgconfluent", "y", "guahwdpdvzealjrnpgiqumxtbqq",
            com.azure.core.util.Context.NONE);
    }
}
