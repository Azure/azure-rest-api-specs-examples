Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Services EnableTestEndpoint. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Services_EnableTestEndpoint.json
     */
    /**
     * Sample code: Services_EnableTestEndpoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void servicesEnableTestEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getServices()
            .enableTestEndpointWithResponse("myResourceGroup", "myservice", Context.NONE);
    }
}
```
