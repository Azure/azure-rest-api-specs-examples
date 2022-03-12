Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Endpoints Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_Start.json
     */
    /**
     * Sample code: Endpoints_Start.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsStart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getEndpoints().start("RG", "profile1", "endpoint1", Context.NONE);
    }
}
```
