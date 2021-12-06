Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.BackendAddressPoolInner;
import com.azure.resourcemanager.network.models.LoadBalancerBackendAddress;
import java.util.Arrays;

/** Samples for LoadBalancerBackendAddressPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LBBackendAddressPoolWithBackendAddressesPut.json
     */
    /**
     * Sample code: Update load balancer backend pool with backend addresses containing virtual network and IP address.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateLoadBalancerBackendPoolWithBackendAddressesContainingVirtualNetworkAndIPAddress(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLoadBalancerBackendAddressPools()
            .createOrUpdate(
                "testrg",
                "lb",
                "backend",
                new BackendAddressPoolInner()
                    .withLoadBalancerBackendAddresses(
                        Arrays
                            .asList(
                                new LoadBalancerBackendAddress()
                                    .withName("address1")
                                    .withVirtualNetwork(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"))
                                    .withIpAddress("10.0.0.4"),
                                new LoadBalancerBackendAddress()
                                    .withName("address2")
                                    .withVirtualNetwork(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"))
                                    .withIpAddress("10.0.0.5"))),
                Context.NONE);
    }
}
```
