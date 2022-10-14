import com.azure.core.util.Context;

/** Samples for GitHubConnectorStatsOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorStatsGet.json
     */
    /**
     * Sample code: GitHubConnectorStats_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubConnectorStatsGet(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubConnectorStatsOperations().getWithResponse("westusrg", "testconnector", Context.NONE);
    }
}
