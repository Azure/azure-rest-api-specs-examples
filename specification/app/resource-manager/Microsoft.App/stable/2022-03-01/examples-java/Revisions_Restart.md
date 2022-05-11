Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.2/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ContainerAppsRevisions RestartRevision. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/Revisions_Restart.json
     */
    /**
     * Sample code: Restart Container App's revision.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void restartContainerAppSRevision(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsRevisions()
            .restartRevisionWithResponse("rg", "testStaticSite0", "testcontainerApp0-pjxhsye", Context.NONE);
    }
}
```
