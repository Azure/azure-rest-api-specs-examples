
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
     * x-ms-original-file: 2025-07-01/NetworkInterfaceCreate.json
     */
    /**
     * Sample code: Create network interface.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createNetworkInterface(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkInterfaces().createOrUpdate("rg1", "test-nic", new NetworkInterfaceInner()
            .withLocation("eastus")
            .withIpConfigurations(Arrays.asList(new NetworkInterfaceIpConfigurationInner().withName("ipconfig1")
                .withSubnet(new SubnetInner().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"))
                .withPublicIpAddress(new PublicIpAddressInner().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip")),
                new NetworkInterfaceIpConfigurationInner().withName("ipconfig2").withPrivateIPAddressPrefixLength(28)))
            .withEnableAcceleratedNetworking(true).withDisableTcpStateTracking(true), com.azure.core.util.Context.NONE);
    }
}
