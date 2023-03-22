/** Samples for SecureScoreControlDefinitions ListBySubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScoreControlDefinitions/ListSecureScoreControlDefinitions_subscription_example.json
     */
    /**
     * Sample code: List security controls definition by subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecurityControlsDefinitionBySubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.secureScoreControlDefinitions().listBySubscription(com.azure.core.util.Context.NONE);
    }
}
