/** Samples for SecurityContacts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/SecurityContacts/DeleteSecurityContact_example.json
     */
    /**
     * Sample code: Deletes a security contact data.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deletesASecurityContactData(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityContacts().deleteWithResponse("default", com.azure.core.util.Context.NONE);
    }
}
