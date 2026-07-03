
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
     * x-ms-original-file: 2025-07-01/NetworkInterfaceCreateGatewayLoadBalancerConsumer.json
     */
    /**
     * Sample code: Create network interface with Gateway Load Balancer Consumer configured.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createNetworkInterfaceWithGatewayLoadBalancerConsumerConfigured(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().createOrUpdate("rg1", "test-nic", new NetworkInterfaceInner()
            .withLocation("eastus")
            .withIpConfigurations(Arrays.asList(new NetworkInterfaceIpConfigurationInner().withName("ipconfig1")
                .withGatewayLoadBalancer(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb-provider"))
                .withSubnet(new SubnetInner().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"))
                .withPublicIpAddress(new PublicIpAddressInner().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"))))
            .withEnableAcceleratedNetworking(true), com.azure.core.util.Context.NONE);
    }
}
