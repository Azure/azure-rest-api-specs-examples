Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.IpConfigurationProfileInner;
import com.azure.resourcemanager.network.fluent.models.NetworkProfileInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.ContainerNetworkInterfaceConfiguration;
import java.util.Arrays;

/** Samples for NetworkProfiles CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkProfileCreateConfigOnly.json
     */
    /**
     * Sample code: Create network profile defaults.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkProfileDefaults(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkProfiles()
            .createOrUpdateWithResponse(
                "rg1",
                "networkProfile1",
                new NetworkProfileInner()
                    .withLocation("westus")
                    .withContainerNetworkInterfaceConfigurations(
                        Arrays
                            .asList(
                                new ContainerNetworkInterfaceConfiguration()
                                    .withName("eth1")
                                    .withIpConfigurations(
                                        Arrays
                                            .asList(
                                                new IpConfigurationProfileInner()
                                                    .withName("ipconfig1")
                                                    .withSubnet(
                                                        new SubnetInner()
                                                            .withId(
                                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1")))))),
                Context.NONE);
    }
}
```
