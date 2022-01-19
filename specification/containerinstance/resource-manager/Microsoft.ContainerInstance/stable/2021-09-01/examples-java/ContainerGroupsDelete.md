Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ContainerGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsDelete.json
     */
    /**
     * Sample code: ContainerGroupsDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().delete("demo", "demo1", Context.NONE);
    }
}
```
