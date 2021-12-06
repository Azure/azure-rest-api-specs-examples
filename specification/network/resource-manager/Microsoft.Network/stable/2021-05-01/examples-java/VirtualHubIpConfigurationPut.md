Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.HubIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;

/** Samples for VirtualHubIpConfiguration CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualHubIpConfigurationPut.json
     */
    /**
     * Sample code: VirtualHubIpConfigurationPut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualHubIpConfigurationPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualHubIpConfigurations()
            .createOrUpdate(
                "rg1",
                "hub1",
                "ipconfig1",
                new HubIpConfigurationInner()
                    .withSubnet(
                        new SubnetInner()
                            .withId(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1")),
                Context.NONE);
    }
}
```
