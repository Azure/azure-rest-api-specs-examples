Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.LoadBalancerVipSwapRequest;
import com.azure.resourcemanager.network.models.LoadBalancerVipSwapRequestFrontendIpConfiguration;
import java.util.Arrays;

/** Samples for LoadBalancers SwapPublicIpAddresses. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancersSwapPublicIpAddresses.json
     */
    /**
     * Sample code: Swap VIPs between two load balancers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void swapVIPsBetweenTwoLoadBalancers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLoadBalancers()
            .swapPublicIpAddresses(
                "westus",
                new LoadBalancerVipSwapRequest()
                    .withFrontendIpConfigurations(
                        Arrays
                            .asList(
                                new LoadBalancerVipSwapRequestFrontendIpConfiguration()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/lbfe1")
                                    .withPublicIpAddress(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPAddresses/pip2")),
                                new LoadBalancerVipSwapRequestFrontendIpConfiguration()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/loadBalancers/lb2/frontendIPConfigurations/lbfe2")
                                    .withPublicIpAddress(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pip1")))),
                Context.NONE);
    }
}
```
