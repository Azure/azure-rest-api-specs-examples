
/**
 * Samples for DevOpsConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/DeleteDevOpsConfigurations_example.json
     */
    /**
     * Sample code: Delete_DevOpsConfigurations.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteDevOpsConfigurations(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsConfigurations().delete("myRg", "mySecurityConnectorName", com.azure.core.util.Context.NONE);
    }
}
