```java
import com.azure.core.util.Context;

/** Samples for ContainerGroups Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsStart.json
     */
    /**
     * Sample code: ContainerStart.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerStart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerGroups().manager().serviceClient().getContainerGroups().start("demo", "demo1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
