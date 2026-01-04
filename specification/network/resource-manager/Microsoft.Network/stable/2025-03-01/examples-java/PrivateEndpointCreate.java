
import com.azure.resourcemanager.network.fluent.models.PrivateEndpointInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.PrivateEndpointIpConfiguration;
import com.azure.resourcemanager.network.models.PrivateEndpointIpVersionType;
import com.azure.resourcemanager.network.models.PrivateLinkServiceConnection;
import java.util.Arrays;

/**
 * Samples for PrivateEndpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PrivateEndpointCreate.json
     */
    /**
     * Sample code: Create private endpoint.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPrivateEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateEndpoints().createOrUpdate("rg1", "testPe",
            new PrivateEndpointInner().withLocation("eastus2euap").withSubnet(new SubnetInner().withId(
                "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"))
                .withIpVersionType(PrivateEndpointIpVersionType.IPV4)
                .withPrivateLinkServiceConnections(Arrays.asList(new PrivateLinkServiceConnection()
                    .withPrivateLinkServiceId(
                        "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls")
                    .withGroupIds(Arrays.asList("groupIdFromResource"))
                    .withRequestMessage("Please approve my connection.")))
                .withIpConfigurations(Arrays.asList(new PrivateEndpointIpConfiguration().withName("pestaticconfig")
                    .withGroupId("file").withMemberName("file").withPrivateIpAddress("192.168.0.6")))
                .withCustomNetworkInterfaceName("testPeNic"),
            com.azure.core.util.Context.NONE);
    }
}
