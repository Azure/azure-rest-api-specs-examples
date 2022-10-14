import com.azure.core.util.Context;

/** Samples for GitHubRepo ListByConnector. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoListByConnector.json
     */
    /**
     * Sample code: GitHubRepo_ListByConnector.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubRepoListByConnector(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubRepoes().listByConnector("westusrg", "testconnector", Context.NONE);
    }
}
