import com.azure.core.util.Context;

/** Samples for SecurityContacts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/SecurityContacts/GetSecurityContact_example.json
     */
    /**
     * Sample code: Get a security contact.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getASecurityContact(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityContacts().getWithResponse("default", Context.NONE);
    }
}
