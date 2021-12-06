Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkInner;
import com.azure.resourcemanager.network.models.AddressSpace;
import java.util.Arrays;

/** Samples for VirtualNetworks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkCreate.json
     */
    /**
     * Sample code: Create virtual network.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualNetwork(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworks()
            .createOrUpdate(
                "rg1",
                "test-vnet",
                new VirtualNetworkInner()
                    .withLocation("eastus")
                    .withAddressSpace(new AddressSpace().withAddressPrefixes(Arrays.asList("10.0.0.0/16")))
                    .withFlowTimeoutInMinutes(10),
                Context.NONE);
    }
}
```
