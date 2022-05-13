Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.2/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ManagedEnvironments GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ManagedEnvironments_Get.json
     */
    /**
     * Sample code: Get environments by name.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getEnvironmentsByName(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().getByResourceGroupWithResponse("examplerg", "jlaw-demo1", Context.NONE);
    }
}
```
