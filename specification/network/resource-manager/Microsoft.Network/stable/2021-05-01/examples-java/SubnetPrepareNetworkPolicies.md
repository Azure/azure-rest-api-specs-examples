Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.PrepareNetworkPoliciesRequest;

/** Samples for Subnets PrepareNetworkPolicies. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/SubnetPrepareNetworkPolicies.json
     */
    /**
     * Sample code: Prepare Network Policies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void prepareNetworkPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSubnets()
            .prepareNetworkPolicies(
                "rg1",
                "test-vnet",
                "subnet1",
                new PrepareNetworkPoliciesRequest().withServiceName("Microsoft.Sql/managedInstances"),
                Context.NONE);
    }
}
```
