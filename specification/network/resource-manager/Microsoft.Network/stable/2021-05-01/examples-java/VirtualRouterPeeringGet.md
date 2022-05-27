Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualRouterPeerings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualRouterPeeringGet.json
     */
    /**
     * Sample code: Get Virtual Router Peering.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualRouterPeering(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualRouterPeerings()
            .getWithResponse("rg1", "virtualRouter", "peering1", Context.NONE);
    }
}
```
