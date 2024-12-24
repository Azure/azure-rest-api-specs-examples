
/**
 * Samples for DevOpsConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/GetDevOpsConfigurationsWithCapabilities_example.json
     */
    /**
     * Sample code: Get_DevOpsConfigurations_WithCapabilities.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getDevOpsConfigurationsWithCapabilities(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsConfigurations().getWithResponse("myRg", "mySecurityConnectorName",
            com.azure.core.util.Context.NONE);
    }
}
