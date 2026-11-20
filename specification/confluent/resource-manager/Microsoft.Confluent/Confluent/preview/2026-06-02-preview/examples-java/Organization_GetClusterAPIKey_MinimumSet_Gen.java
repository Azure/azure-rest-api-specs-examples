
/**
 * Samples for Organization GetClusterApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/Organization_GetClusterAPIKey_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organization_GetClusterAPIKey_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationGetClusterAPIKeyMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getClusterApiKeyWithResponse("rgconfluent", "oiywgdcgyrmdcquutyn",
            "gmgzzzwsoctmbdrgttw", com.azure.core.util.Context.NONE);
    }
}
