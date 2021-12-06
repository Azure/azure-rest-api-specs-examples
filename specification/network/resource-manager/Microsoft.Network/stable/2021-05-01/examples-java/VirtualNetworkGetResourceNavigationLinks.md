Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ResourceNavigationLinks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGetResourceNavigationLinks.json
     */
    /**
     * Sample code: Get Resource Navigation Links.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getResourceNavigationLinks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getResourceNavigationLinks()
            .listWithResponse("rg1", "vnet", "subnet", Context.NONE);
    }
}
```
