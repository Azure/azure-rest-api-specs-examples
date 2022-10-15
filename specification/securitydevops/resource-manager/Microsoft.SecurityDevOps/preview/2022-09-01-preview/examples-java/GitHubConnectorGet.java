import com.azure.core.util.Context;

/** Samples for GitHubConnector GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorGet.json
     */
    /**
     * Sample code: GitHubConnector_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubConnectorGet(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.gitHubConnectors().getByResourceGroupWithResponse("westusrg", "testconnector", Context.NONE);
    }
}
