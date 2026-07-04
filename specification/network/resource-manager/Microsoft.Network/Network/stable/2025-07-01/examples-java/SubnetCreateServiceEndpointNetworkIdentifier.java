
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.ServiceEndpointPropertiesFormat;
import java.util.Arrays;

/**
 * Samples for Subnets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubnetCreateServiceEndpointNetworkIdentifier.json
     */
    /**
     * Sample code: Create subnet with service endpoints with network identifier.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createSubnetWithServiceEndpointsWithNetworkIdentifier(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubnets().createOrUpdate("subnet-test", "vnetname", "subnet1", new SubnetInner()
            .withAddressPrefix("10.0.0.0/16")
            .withServiceEndpoints(Arrays.asList(new ServiceEndpointPropertiesFormat().withService("Microsoft.Storage")
                .withNetworkIdentifier(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/subnet-test/providers/Microsoft.Network/publicIPAddresses/test-ip")))),
            com.azure.core.util.Context.NONE);
    }
}
