import com.azure.core.util.Context;

/** Samples for GitHubConnector Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorDelete.json
     */
    /**
     * Sample code: GitHubConnector_Delete.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubConnectorDelete(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubConnectors().delete("westusrg", "testconnector", Context.NONE);
    }
}
