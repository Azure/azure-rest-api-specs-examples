
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkInner;
import com.azure.resourcemanager.network.models.AddressSpace;
import com.azure.resourcemanager.network.models.ServiceEndpointPropertiesFormat;
import java.util.Arrays;

/**
 * Samples for VirtualNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * VirtualNetworkCreateServiceEndpoints.json
     */
    /**
     * Sample code: Create virtual network with service endpoints.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualNetworkWithServiceEndpoints(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworks().createOrUpdate("vnetTest", "vnet1",
            new VirtualNetworkInner().withLocation("eastus")
                .withAddressSpace(new AddressSpace().withAddressPrefixes(Arrays.asList("10.0.0.0/16")))
                .withSubnets(Arrays
                    .asList(new SubnetInner().withName("test-1").withAddressPrefix("10.0.0.0/16").withServiceEndpoints(
                        Arrays.asList(new ServiceEndpointPropertiesFormat().withService("Microsoft.Storage"))))),
            com.azure.core.util.Context.NONE);
    }
}
