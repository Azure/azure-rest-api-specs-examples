Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ContainerServices ListOrchestrators. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2019-08-01/examples/ContainerServiceListOrchestrators.json
     */
    /**
     * Sample code: List Container Service Orchestrators.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainerServiceOrchestrators(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getContainerServices()
            .listOrchestratorsWithResponse("location1", null, Context.NONE);
    }
}
```
