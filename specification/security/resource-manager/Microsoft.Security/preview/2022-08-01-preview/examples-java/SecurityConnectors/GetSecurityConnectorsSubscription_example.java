import com.azure.core.util.Context;

/** Samples for SecurityConnectors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/GetSecurityConnectorsSubscription_example.json
     */
    /**
     * Sample code: List all security connectors of a specified subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listAllSecurityConnectorsOfASpecifiedSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityConnectors().list(Context.NONE);
    }
}
