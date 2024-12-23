
/**
 * Samples for AzureDevOpsOrgs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/GetAzureDevOpsOrgs_example.json
     */
    /**
     * Sample code: Get_AzureDevOpsOrgs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getAzureDevOpsOrgs(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsOrgs().getWithResponse("myRg", "mySecurityConnectorName", "myAzDevOpsOrg",
            com.azure.core.util.Context.NONE);
    }
}
