
/**
 * Samples for SecureScoreControlDefinitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/
     * secureScoreControlDefinitions/ListSecureScoreControlDefinitions_example.json
     */
    /**
     * Sample code: List security controls definition.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecurityControlsDefinition(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.secureScoreControlDefinitions().list(com.azure.core.util.Context.NONE);
    }
}
