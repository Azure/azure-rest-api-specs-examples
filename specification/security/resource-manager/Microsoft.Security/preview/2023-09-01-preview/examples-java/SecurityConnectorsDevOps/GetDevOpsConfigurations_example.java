
/**
 * Samples for DevOpsConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/GetDevOpsConfigurations_example.json
     */
    /**
     * Sample code: Get_DevOpsConfigurations.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getDevOpsConfigurations(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsConfigurations().getWithResponse("myRg", "mySecurityConnectorName",
            com.azure.core.util.Context.NONE);
    }
}
