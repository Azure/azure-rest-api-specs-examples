Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Features List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/listProviderFeatures.json
     */
    /**
     * Sample code: List provider Features.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProviderFeatures(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .featureClient()
            .getFeatures()
            .list("Resource Provider Namespace", Context.NONE);
    }
}
```
