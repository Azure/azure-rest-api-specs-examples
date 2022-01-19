Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AgentPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/AgentPoolsGet.json
     */
    /**
     * Sample code: AgentPools_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void agentPoolsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getAgentPools()
            .getWithResponse("myResourceGroup", "myRegistry", "myAgentPool", Context.NONE);
    }
}
```
