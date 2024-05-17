
/**
 * Samples for GitLabProjects Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * GetGitLabProjects_example.json
     */
    /**
     * Sample code: Get_GitLabProjects.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getGitLabProjects(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitLabProjects().getWithResponse("myRg", "mySecurityConnectorName", "myGitLabGroup$mySubGroup",
            "myGitLabProject", com.azure.core.util.Context.NONE);
    }
}
