
/**
 * Samples for AzureDevOpsRepos List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/ListAzureDevOpsRepos_example.json
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
