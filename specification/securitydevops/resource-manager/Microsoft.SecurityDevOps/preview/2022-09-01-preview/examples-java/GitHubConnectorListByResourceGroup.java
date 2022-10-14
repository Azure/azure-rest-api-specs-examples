import com.azure.core.util.Context;

/** Samples for GitHubConnector ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorListByResourceGroup.json
     */
    /**
     * Sample code: GitHubConnector_ListByResourceGroup.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubConnectorListByResourceGroup(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubConnectors().listByResourceGroup("westusrg", Context.NONE);
    }
}
