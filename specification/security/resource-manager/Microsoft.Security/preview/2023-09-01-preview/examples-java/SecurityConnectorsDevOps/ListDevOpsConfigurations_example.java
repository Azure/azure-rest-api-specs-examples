
/**
 * Samples for DevOpsConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/ListDevOpsConfigurations_example.json
     */
    /**
     * Sample code: List_DevOpsConfigurations.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listDevOpsConfigurations(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsConfigurations().list("myRg", "mySecurityConnectorName", com.azure.core.util.Context.NONE);
    }
}
