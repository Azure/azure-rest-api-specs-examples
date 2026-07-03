
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkInner;
import com.azure.resourcemanager.network.models.AddressSpace;
import com.azure.resourcemanager.network.models.VirtualNetworkBgpCommunities;
import java.util.Arrays;

/**
 * Samples for VirtualNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkCreateWithBgpCommunities.json
     */
    /**
     * Sample code: Create virtual network with Bgp Communities.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createVirtualNetworkWithBgpCommunities(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().createOrUpdate("rg1", "test-vnet",
            new VirtualNetworkInner().withLocation("eastus")
                .withAddressSpace(new AddressSpace().withAddressPrefixes(Arrays.asList("10.0.0.0/16")))
                .withSubnets(Arrays.asList(new SubnetInner().withName("test-1").withAddressPrefix("10.0.0.0/24")))
                .withBgpCommunities(new VirtualNetworkBgpCommunities().withVirtualNetworkCommunity("12076:20000")),
            com.azure.core.util.Context.NONE);
    }
}
