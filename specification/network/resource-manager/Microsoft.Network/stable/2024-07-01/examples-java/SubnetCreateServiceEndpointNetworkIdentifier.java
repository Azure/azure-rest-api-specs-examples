
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.ServiceEndpointPropertiesFormat;
import java.util.Arrays;

/**
 * Samples for Subnets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * SubnetCreateServiceEndpointNetworkIdentifier.json
     */
    /**
     * Sample code: Create subnet with service endpoints with network identifier.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createSubnetWithServiceEndpointsWithNetworkIdentifier(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSubnets().createOrUpdate("subnet-test", "vnetname", "subnet1",
            new SubnetInner().withAddressPrefix("10.0.0.0/16")
                .withServiceEndpoints(Arrays.asList(new ServiceEndpointPropertiesFormat()
                    .withService("Microsoft.Storage")
                    .withNetworkIdentifier(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/subnet-test/providers/Microsoft.Network/publicIPAddresses/test-ip")))),
            com.azure.core.util.Context.NONE);
    }
}
