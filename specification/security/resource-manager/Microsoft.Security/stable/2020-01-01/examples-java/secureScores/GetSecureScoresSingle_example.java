
/**
 * Samples for SecureScores Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/
     * GetSecureScoresSingle_example.json
     */
    /**
     * Sample code: Get single secure score.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSingleSecureScore(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.secureScores().getWithResponse("ascScore", com.azure.core.util.Context.NONE);
    }
}
