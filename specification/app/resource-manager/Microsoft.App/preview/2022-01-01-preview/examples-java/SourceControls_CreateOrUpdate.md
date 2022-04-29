Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.1/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.appcontainers.models.AzureCredentials;
import com.azure.resourcemanager.appcontainers.models.GithubActionConfiguration;
import com.azure.resourcemanager.appcontainers.models.RegistryInfo;

/** Samples for ContainerAppsSourceControls CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/SourceControls_CreateOrUpdate.json
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
                            .withTenantId("<tenantid>")))
            .create();
    }
}
```
