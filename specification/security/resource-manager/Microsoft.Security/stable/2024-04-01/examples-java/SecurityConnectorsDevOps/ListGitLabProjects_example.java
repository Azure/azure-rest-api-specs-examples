
/**
 * Samples for GitLabProjects List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * ListGitLabProjects_example.json
     */
    /**
     * Sample code: List_GitLabProjects.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listGitLabProjects(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitLabProjects().list("myRg", "mySecurityConnectorName", "myGitLabGroup$mySubGroup",
            com.azure.core.util.Context.NONE);
    }
}
