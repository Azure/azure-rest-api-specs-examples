Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Bindings Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Bindings_Delete.json
     */
    /**
     * Sample code: Bindings_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void bindingsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBindings()
            .delete("myResourceGroup", "myservice", "myapp", "mybinding", Context.NONE);
    }
}
```
