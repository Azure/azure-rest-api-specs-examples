
/**
 * Samples for SecurityContacts List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-12-01-preview/examples/SecurityContacts/
     * GetSecurityContactsSubscription_example.json
     */
    /**
     * Sample code: List security contact data.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecurityContactData(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityContacts().list(com.azure.core.util.Context.NONE);
    }
}
