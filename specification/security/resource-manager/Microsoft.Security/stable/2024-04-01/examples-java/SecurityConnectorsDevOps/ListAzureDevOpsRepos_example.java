
/**
 * Samples for AzureDevOpsRepos List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * ListAzureDevOpsRepos_example.json
     */
    /**
     * Sample code: List_AzureDevOpsRepos.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listAzureDevOpsRepos(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsRepos().list("myRg", "mySecurityConnectorName", "myAzDevOpsOrg", "myAzDevOpsProject",
            com.azure.core.util.Context.NONE);
    }
}
