import com.azure.core.util.Context;

/** Samples for SecurityConnectors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/DeleteSecurityConnector_example.json
     */
    /**
     * Sample code: Delete a security connector.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteASecurityConnector(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityConnectors().deleteWithResponse("myRg", "mySecurityConnectorName", Context.NONE);
    }
}
