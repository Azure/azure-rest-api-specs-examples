Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.2/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DaprComponents ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/DaprComponents_ListSecrets.json
     */
    /**
     * Sample code: List Container Apps Secrets.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppsSecrets(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponents().listSecretsWithResponse("examplerg", "myenvironment", "reddog", Context.NONE);
    }
}
```
