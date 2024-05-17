
/**
 * Samples for GitHubRepos List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * ListGitHubRepos_example.json
     */
    /**
     * Sample code: List_GitHubRepos.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listGitHubRepos(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitHubRepos().list("myRg", "mySecurityConnectorName", "myGitHubOwner",
            com.azure.core.util.Context.NONE);
    }
}
