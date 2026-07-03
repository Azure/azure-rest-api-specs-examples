
import com.azure.resourcemanager.network.fluent.models.PrivateEndpointInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.PrivateEndpointIpConfiguration;
import com.azure.resourcemanager.network.models.PrivateLinkServiceConnection;
import java.util.Arrays;

/**
 * Samples for PrivateEndpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateEndpointCreateForManualApproval.json
     */
    /**
     * Sample code: Create private endpoint with manual approval connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createPrivateEndpointWithManualApprovalConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateEndpoints().createOrUpdate("rg1", "testPe", new PrivateEndpointInner()
            .withLocation("eastus")
            .withSubnet(new SubnetInner().withId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"))
            .withManualPrivateLinkServiceConnections(Arrays.asList(new PrivateLinkServiceConnection()
                .withPrivateLinkServiceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls")
                .withGroupIds(Arrays.asList("groupIdFromResource"))
                .withRequestMessage("Please manually approve my connection.")))
            .withIpConfigurations(Arrays.asList(new PrivateEndpointIpConfiguration().withName("pestaticconfig")
                .withGroupId("file").withMemberName("file").withPrivateIpAddress("192.168.0.5")))
            .withCustomNetworkInterfaceName("testPeNic"), com.azure.core.util.Context.NONE);
    }
}
