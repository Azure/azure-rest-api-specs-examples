Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Location ListUsage. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupUsage.json
     */
    /**
     * Sample code: ContainerUsage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getLocations().listUsage("westcentralus", Context.NONE);
    }
}
```
