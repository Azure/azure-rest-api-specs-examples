/** Samples for GitHubOwner CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubOwnerCreateOrUpdate.json
     */
    /**
     * Sample code: GitHubOwner_CreateOrUpdate.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubOwnerCreateOrUpdate(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubOwners().define("Azure").withExistingGitHubConnector("westusrg", "testconnector").create();
    }
}
