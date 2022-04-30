Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.1/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ContainerAppsRevisionReplicas ListReplicas. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/Replicas_List.json
     */
    /**
     * Sample code: List Container App's replicas.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listContainerAppSReplicas(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsRevisionReplicas()
            .listReplicasWithResponse("workerapps-rg-xj", "myapp", "myapp--0wlqy09", Context.NONE);
    }
}
```
