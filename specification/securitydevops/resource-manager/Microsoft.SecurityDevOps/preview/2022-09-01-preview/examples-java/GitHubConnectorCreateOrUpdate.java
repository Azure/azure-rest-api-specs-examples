import com.azure.resourcemanager.securitydevops.models.GitHubConnectorProperties;

/** Samples for GitHubConnector CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorCreateOrUpdate.json
     */
    /**
     * Sample code: GitHubConnector_CreateOrUpdate.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void gitHubConnectorCreateOrUpdate(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager
            .gitHubConnectors()
            .define("testconnector")
            .withRegion("West US")
            .withExistingResourceGroup("westusrg")
            .withProperties(new GitHubConnectorProperties().withCode("00000000000000000000"))
            .create();
    }
}
