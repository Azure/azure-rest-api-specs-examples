
/**
 * Samples for SecureScoreControls List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/
     * ListSecureScoreControls_example.json
     */
    /**
     * Sample code: List all secure scores controls.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAllSecureScoresControls(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.secureScoreControls().list(null, com.azure.core.util.Context.NONE);
    }
}
