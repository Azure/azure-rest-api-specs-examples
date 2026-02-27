
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.NetworkInterfaceInner;
import com.azure.resourcemanager.network.fluent.models.NetworkInterfaceIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import java.util.Arrays;

/**
 * Samples for NetworkInterfaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkInterfaceCreateGatewayLoadBalancerConsumer.json
     */
    /**
     * Sample code: Create network interface with Gateway Load Balancer Consumer configured.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkInterfaceWithGatewayLoadBalancerConsumerConfigured(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces().createOrUpdate("rg1", "test-nic",
            new NetworkInterfaceInner().withLocation("eastus")
                .withIpConfigurations(Arrays.asList(new NetworkInterfaceIpConfigurationInner().withName("ipconfig1")
                    .withGatewayLoadBalancer(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb-provider"))
                    .withSubnet(new SubnetInner().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"))
                    .withPublicIpAddress(new PublicIpAddressInner().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"))))
                .withEnableAcceleratedNetworking(true),
            com.azure.core.util.Context.NONE);
    }
}
