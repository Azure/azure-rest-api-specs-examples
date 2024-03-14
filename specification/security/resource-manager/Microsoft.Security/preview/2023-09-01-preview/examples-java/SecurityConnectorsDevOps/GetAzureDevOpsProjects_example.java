
/**
 * Samples for AzureDevOpsProjects Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/GetAzureDevOpsProjects_example.json
     */
    /**
     * Sample code: Get_AzureDevOpsProjects.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getAzureDevOpsProjects(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsProjects().getWithResponse("myRg", "mySecurityConnectorName", "myAzDevOpsOrg",
            "myAzDevOpsProject", com.azure.core.util.Context.NONE);
    }
}
