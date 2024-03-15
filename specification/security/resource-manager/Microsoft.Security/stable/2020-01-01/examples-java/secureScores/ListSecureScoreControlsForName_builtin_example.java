
/**
 * Samples for SecureScoreControls ListBySecureScore.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/
     * ListSecureScoreControlsForName_builtin_example.json
     */
    /**
     * Sample code: Get security controls and their current score for the specified initiative.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityControlsAndTheirCurrentScoreForTheSpecifiedInitiative(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.secureScoreControls().listBySecureScore("ascScore", null, com.azure.core.util.Context.NONE);
    }
}
