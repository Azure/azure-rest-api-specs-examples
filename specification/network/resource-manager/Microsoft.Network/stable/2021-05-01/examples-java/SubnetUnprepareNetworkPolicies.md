Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.UnprepareNetworkPoliciesRequest;

/** Samples for Subnets UnprepareNetworkPolicies. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/SubnetUnprepareNetworkPolicies.json
     */
    /**
     * Sample code: Unprepare Network Policies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void unprepareNetworkPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSubnets()
            .unprepareNetworkPolicies(
                "rg1",
                "test-vnet",
                "subnet1",
                new UnprepareNetworkPoliciesRequest().withServiceName("Microsoft.Sql/managedInstances"),
                Context.NONE);
    }
}
```
