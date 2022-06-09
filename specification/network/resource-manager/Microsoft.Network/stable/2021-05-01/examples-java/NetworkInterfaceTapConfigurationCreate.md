```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.NetworkInterfaceTapConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkTapInner;

/** Samples for NetworkInterfaceTapConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkInterfaceTapConfigurationCreate.json
     */
    /**
     * Sample code: Create Network Interface Tap Configurations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkInterfaceTapConfigurations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaceTapConfigurations()
            .createOrUpdate(
                "testrg",
                "mynic",
                "tapconfiguration1",
                new NetworkInterfaceTapConfigurationInner()
                    .withVirtualNetworkTap(
                        new VirtualNetworkTapInner()
                            .withId(
                                "/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworkTaps/testvtap")),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
