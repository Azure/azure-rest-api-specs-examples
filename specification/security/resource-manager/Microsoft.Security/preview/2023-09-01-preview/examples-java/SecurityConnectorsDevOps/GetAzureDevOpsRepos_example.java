
/**
 * Samples for AzureDevOpsRepos Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/GetAzureDevOpsRepos_example.json
     */
    /**
     * Sample code: Get_AzureDevOpsRepos.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getAzureDevOpsRepos(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsRepos().getWithResponse("myRg", "mySecurityConnectorName", "myAzDevOpsOrg",
            "myAzDevOpsProject", "myAzDevOpsRepo", com.azure.core.util.Context.NONE);
    }
}
