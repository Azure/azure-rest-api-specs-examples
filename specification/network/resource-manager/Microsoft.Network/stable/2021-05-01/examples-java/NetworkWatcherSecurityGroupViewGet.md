Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.SecurityGroupViewParameters;

/** Samples for NetworkWatchers GetVMSecurityRules. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherSecurityGroupViewGet.json
     */
    /**
     * Sample code: Get security group view.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSecurityGroupView(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .getVMSecurityRules(
                "rg1",
                "nw1",
                new SecurityGroupViewParameters()
                    .withTargetResourceId(
                        "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
                Context.NONE);
    }
}
```
