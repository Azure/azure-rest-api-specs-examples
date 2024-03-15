
/**
 * Samples for GitLabGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/
     * SecurityConnectorsDevOps/GetGitLabGroups_example.json
     */
    /**
     * Sample code: Get_GitLabGroups.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getGitLabGroups(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.gitLabGroups().getWithResponse("myRg", "mySecurityConnectorName", "myGitLabGroup$mySubGroup",
            com.azure.core.util.Context.NONE);
    }
}
