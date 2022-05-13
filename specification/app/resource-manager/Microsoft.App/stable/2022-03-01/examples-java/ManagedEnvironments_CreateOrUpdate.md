Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.2/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.appcontainers.models.AppLogsConfiguration;
import com.azure.resourcemanager.appcontainers.models.LogAnalyticsConfiguration;

/** Samples for ManagedEnvironments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ManagedEnvironments_CreateOrUpdate.json
     */
    /**
     * Sample code: Create environments.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createEnvironments(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .managedEnvironments()
            .define("testcontainerenv")
            .withRegion("East US")
            .withExistingResourceGroup("examplerg")
            .withDaprAIConnectionString(
                "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/")
            .withAppLogsConfiguration(
                new AppLogsConfiguration()
                    .withLogAnalyticsConfiguration(
                        new LogAnalyticsConfiguration().withCustomerId("string").withSharedKey("string")))
            .withZoneRedundant(true)
            .create();
    }
}
```
