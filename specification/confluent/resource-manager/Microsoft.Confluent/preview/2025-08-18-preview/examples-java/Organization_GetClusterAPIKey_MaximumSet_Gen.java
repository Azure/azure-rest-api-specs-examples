
/**
 * Samples for Organization GetClusterApiKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Organization_GetClusterAPIKey_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_GetClusterAPIKey_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationGetClusterAPIKeyMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().getClusterApiKeyWithResponse("rgconfluent", "puauqgrwsfgmolfhazfjcavnj",
            "xxsquwnsllkkzuyzlhdxdl", com.azure.core.util.Context.NONE);
    }
}
