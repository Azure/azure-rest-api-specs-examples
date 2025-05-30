
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.models.LoadBalancerVipSwapRequest;
import com.azure.resourcemanager.network.models.LoadBalancerVipSwapRequestFrontendIpConfiguration;
import java.util.Arrays;

/**
 * Samples for LoadBalancers SwapPublicIpAddresses.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/
     * LoadBalancersSwapPublicIpAddresses.json
     */
    /**
     * Sample code: Swap VIPs between two load balancers.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void swapVIPsBetweenTwoLoadBalancers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().swapPublicIpAddresses("westus",
            new LoadBalancerVipSwapRequest().withFrontendIpConfigurations(Arrays.asList(
                new LoadBalancerVipSwapRequestFrontendIpConfiguration().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb1/frontendIPConfigurations/lbfe1")
                    .withPublicIpAddress(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/publicIPAddresses/pip2")),
                new LoadBalancerVipSwapRequestFrontendIpConfiguration().withId(
                    "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/loadBalancers/lb2/frontendIPConfigurations/lbfe2")
                    .withPublicIpAddress(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pip1")))),
            com.azure.core.util.Context.NONE);
    }
}
