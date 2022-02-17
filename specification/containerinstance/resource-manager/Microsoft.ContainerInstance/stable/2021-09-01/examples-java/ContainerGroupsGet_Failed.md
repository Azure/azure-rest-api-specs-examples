Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ContainerGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsGet_Failed.json
     */
    /**
     * Sample code: ContainerGroupsGet_Failed.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsGetFailed(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerGroups()
            .manager()
            .serviceClient()
            .getContainerGroups()
            .getByResourceGroupWithResponse("demo", "demo1", Context.NONE);
    }
}
```
