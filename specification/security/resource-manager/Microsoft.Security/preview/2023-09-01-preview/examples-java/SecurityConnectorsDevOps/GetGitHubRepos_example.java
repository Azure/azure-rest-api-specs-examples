
/**
 * Samples for GitHubRepos Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/GetGitHubRepos_example.json
     */
    /**
     * Sample code: Get_GitHubRepos.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getGitHubRepos(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitHubRepos().getWithResponse("myRg", "mySecurityConnectorName", "myGitHubOwner", "myGitHubRepo",
            com.azure.core.util.Context.NONE);
    }
}
