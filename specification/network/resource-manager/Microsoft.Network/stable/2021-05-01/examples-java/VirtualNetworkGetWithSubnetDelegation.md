Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualNetworks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGetWithSubnetDelegation.json
     */
    /**
     * Sample code: Get virtual network with a delegated subnet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkWithADelegatedSubnet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworks()
            .getByResourceGroupWithResponse("rg1", "test-vnet", null, Context.NONE);
    }
}
```
