
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkInner;
import com.azure.resourcemanager.network.models.AddressSpace;
import java.util.Arrays;

/**
 * Samples for VirtualNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkCreateSubnet.
     * json
     */
    /**
     * Sample code: Create virtual network with subnet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualNetworkWithSubnet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworks().createOrUpdate("rg1", "test-vnet",
            new VirtualNetworkInner().withLocation("eastus")
                .withAddressSpace(new AddressSpace().withAddressPrefixes(Arrays.asList("10.0.0.0/16")))
                .withSubnets(Arrays.asList(new SubnetInner().withName("test-1").withAddressPrefix("10.0.0.0/24"))),
            com.azure.core.util.Context.NONE);
    }
}
