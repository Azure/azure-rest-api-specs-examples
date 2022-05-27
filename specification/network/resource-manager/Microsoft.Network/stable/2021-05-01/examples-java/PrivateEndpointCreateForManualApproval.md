Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.PrivateEndpointInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.PrivateEndpointIpConfiguration;
import com.azure.resourcemanager.network.models.PrivateLinkServiceConnection;
import java.util.Arrays;

/** Samples for PrivateEndpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointCreateForManualApproval.json
     */
    /**
     * Sample code: Create private endpoint with manual approval connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPrivateEndpointWithManualApprovalConnection(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateEndpoints()
            .createOrUpdate(
                "rg1",
                "testPe",
                new PrivateEndpointInner()
                    .withLocation("eastus")
                    .withSubnet(
                        new SubnetInner()
                            .withId(
                                "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"))
                    .withManualPrivateLinkServiceConnections(
                        Arrays
                            .asList(
                                new PrivateLinkServiceConnection()
                                    .withPrivateLinkServiceId(
                                        "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls")
                                    .withGroupIds(Arrays.asList("groupIdFromResource"))
                                    .withRequestMessage("Please manually approve my connection.")))
                    .withIpConfigurations(
                        Arrays
                            .asList(
                                new PrivateEndpointIpConfiguration()
                                    .withName("pestaticconfig")
                                    .withGroupId("file")
                                    .withMemberName("file")
                                    .withPrivateIpAddress("192.168.0.5")))
                    .withCustomNetworkInterfaceName("testPeNic"),
                Context.NONE);
    }
}
```
