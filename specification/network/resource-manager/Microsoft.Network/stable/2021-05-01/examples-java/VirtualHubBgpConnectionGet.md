Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualHubBgpConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualHubBgpConnectionGet.json
     */
    /**
     * Sample code: VirtualHubVirtualHubRouteTableV2Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubVirtualHubRouteTableV2Get(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualHubBgpConnections()
            .getWithResponse("rg1", "hub1", "conn1", Context.NONE);
    }
}
```
