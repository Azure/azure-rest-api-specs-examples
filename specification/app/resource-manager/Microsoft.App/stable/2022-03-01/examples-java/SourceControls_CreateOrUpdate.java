import com.azure.resourcemanager.appcontainers.models.AzureCredentials;
import com.azure.resourcemanager.appcontainers.models.GithubActionConfiguration;
import com.azure.resourcemanager.appcontainers.models.RegistryInfo;

/** Samples for ContainerAppsSourceControls CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/SourceControls_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Container App SourceControl.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateContainerAppSourceControl(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsSourceControls()
            .define("current")
            .withExistingContainerApp("workerapps-rg-xj", "testcanadacentral")
            .withRepoUrl("https://github.com/xwang971/ghatest")
            .withBranch("master")
            .withGithubActionConfiguration(
                new GithubActionConfiguration()
                    .withRegistryInfo(
                        new RegistryInfo()
                            .withRegistryUrl("xwang971reg.azurecr.io")
                            .withRegistryUsername("xwang971reg")
                            .withRegistryPassword("<registrypassword>"))
                    .withAzureCredentials(
                        new AzureCredentials()
                            .withClientId("<clientid>")
                            .withClientSecret("<clientsecret>")
                            .withTenantId("<tenantid>"))
                    .withContextPath("./")
                    .withImage("image/tag"))
            .create();
    }
}
