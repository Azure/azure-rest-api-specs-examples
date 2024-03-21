
import com.azure.resourcemanager.appcontainers.models.AzureCredentials;
import com.azure.resourcemanager.appcontainers.models.EnvironmentVariable;
import com.azure.resourcemanager.appcontainers.models.GithubActionConfiguration;
import com.azure.resourcemanager.appcontainers.models.RegistryInfo;
import java.util.Arrays;

/**
 * Samples for ContainerAppsSourceControls CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/
     * SourceControls_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Container App SourceControl.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateContainerAppSourceControl(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSourceControls().define("current")
            .withExistingContainerApp("workerapps-rg-xj", "testcanadacentral")
            .withRepoUrl("https://github.com/xwang971/ghatest").withBranch("master")
            .withGithubActionConfiguration(new GithubActionConfiguration()
                .withRegistryInfo(new RegistryInfo().withRegistryUrl("test-registry.azurecr.io")
                    .withRegistryUsername("test-registry").withRegistryPassword("fakeTokenPlaceholder"))
                .withAzureCredentials(new AzureCredentials().withClientId("<clientid>")
                    .withClientSecret("fakeTokenPlaceholder").withTenantId("<tenantid>").withKind("feaderated"))
                .withContextPath("./").withGithubPersonalAccessToken("fakeTokenPlaceholder").withImage("image/tag")
                .withBuildEnvironmentVariables(
                    Arrays.asList(new EnvironmentVariable().withName("foo1").withValue("bar1"),
                        new EnvironmentVariable().withName("foo2").withValue("bar2"))))
            .create();
    }
}
