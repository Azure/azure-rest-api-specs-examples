Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.NetworkInterfaceIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkTapInner;

/** Samples for VirtualNetworkTaps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkTapCreate.json
     */
    /**
     * Sample code: Create Virtual Network Tap.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualNetworkTap(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkTaps()
            .createOrUpdate(
                "rg1",
                "test-vtap",
                new VirtualNetworkTapInner()
                    .withLocation("centraluseuap")
                    .withDestinationNetworkInterfaceIpConfiguration(
                        new NetworkInterfaceIpConfigurationInner()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/ipconfig1")),
                Context.NONE);
    }
}
```
