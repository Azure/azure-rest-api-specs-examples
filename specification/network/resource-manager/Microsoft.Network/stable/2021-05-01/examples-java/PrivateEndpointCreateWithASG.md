Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ApplicationSecurityGroupInner;
import com.azure.resourcemanager.network.fluent.models.PrivateEndpointInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.PrivateLinkServiceConnection;
import java.util.Arrays;

/** Samples for PrivateEndpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointCreateWithASG.json
     */
    /**
     * Sample code: Create private endpoint with application security groups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createPrivateEndpointWithApplicationSecurityGroups(
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
                    .withLocation("eastus2euap")
                    .withSubnet(
                        new SubnetInner()
                            .withId(
                                "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"))
                    .withPrivateLinkServiceConnections(
                        Arrays
                            .asList(
                                new PrivateLinkServiceConnection()
                                    .withPrivateLinkServiceId(
                                        "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls")
                                    .withGroupIds(Arrays.asList("groupIdFromResource"))
                                    .withRequestMessage("Please approve my connection.")))
                    .withApplicationSecurityGroups(
                        Arrays
                            .asList(
                                new ApplicationSecurityGroupInner()
                                    .withId(
                                        "/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/applicationSecurityGroup/asg1"))),
                Context.NONE);
    }
}
```
