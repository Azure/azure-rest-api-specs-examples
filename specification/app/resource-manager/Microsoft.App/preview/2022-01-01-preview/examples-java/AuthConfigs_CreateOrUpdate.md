Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.1/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.appcontainers.models.AppRegistration;
import com.azure.resourcemanager.appcontainers.models.AuthPlatform;
import com.azure.resourcemanager.appcontainers.models.Facebook;
import com.azure.resourcemanager.appcontainers.models.GlobalValidation;
import com.azure.resourcemanager.appcontainers.models.IdentityProviders;
import com.azure.resourcemanager.appcontainers.models.UnauthenticatedClientActionV2;

/** Samples for ContainerAppsAuthConfigs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/AuthConfigs_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Container App AuthConfig.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateContainerAppAuthConfig(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsAuthConfigs()
            .define("current")
            .withExistingContainerApp("workerapps-rg-xj", "testcanadacentral")
            .withPlatform(new AuthPlatform().withEnabled(true))
            .withGlobalValidation(
                new GlobalValidation().withUnauthenticatedClientAction(UnauthenticatedClientActionV2.ALLOW_ANONYMOUS))
            .withIdentityProviders(
                new IdentityProviders()
                    .withFacebook(
                        new Facebook()
                            .withRegistration(
                                new AppRegistration().withAppId("123").withAppSecretSettingName("facebook-secret"))))
            .create();
    }
}
```
