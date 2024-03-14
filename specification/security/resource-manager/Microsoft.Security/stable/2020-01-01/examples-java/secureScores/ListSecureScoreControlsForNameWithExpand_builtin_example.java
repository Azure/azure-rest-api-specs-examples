
import com.azure.resourcemanager.security.models.ExpandControlsEnum;

/**
 * Samples for SecureScoreControls ListBySecureScore.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScores/
     * ListSecureScoreControlsForNameWithExpand_builtin_example.json
     */
    /**
     * Sample code: Get security controls and their current score for the specified initiative with the expand
     * parameter.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityControlsAndTheirCurrentScoreForTheSpecifiedInitiativeWithTheExpandParameter(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.secureScoreControls().listBySecureScore("ascScore", ExpandControlsEnum.DEFINITION,
            com.azure.core.util.Context.NONE);
    }
}
