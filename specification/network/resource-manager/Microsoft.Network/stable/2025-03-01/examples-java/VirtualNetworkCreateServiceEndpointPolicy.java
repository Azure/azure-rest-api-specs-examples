
import com.azure.resourcemanager.network.fluent.models.ServiceEndpointPolicyInner;
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
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkCreateServiceEndpointPolicy.json
     */
    /**
     * Sample code: Create virtual network with service endpoints and service endpoint policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualNetworkWithServiceEndpointsAndServiceEndpointPolicy(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworks().createOrUpdate("vnetTest", "vnet1",
            new VirtualNetworkInner().withLocation("eastus2euap")
                .withAddressSpace(new AddressSpace().withAddressPrefixes(Arrays.asList("10.0.0.0/16")))
                .withSubnets(Arrays.asList(new SubnetInner().withName("test-1").withAddressPrefix("10.0.0.0/16")
                    .withServiceEndpoints(
                        Arrays.asList(new ServiceEndpointPropertiesFormat().withService("Microsoft.Storage")))
                    .withServiceEndpointPolicies(Arrays.asList(new ServiceEndpointPolicyInner().withId(
                        "/subscriptions/subid/resourceGroups/vnetTest/providers/Microsoft.Network/serviceEndpointPolicies/ServiceEndpointPolicy1"))))),
            com.azure.core.util.Context.NONE);
    }
}
