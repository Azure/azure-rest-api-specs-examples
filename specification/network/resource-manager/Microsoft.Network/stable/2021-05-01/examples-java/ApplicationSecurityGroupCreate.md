Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.ApplicationSecurityGroupInner;

/** Samples for ApplicationSecurityGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationSecurityGroupCreate.json
     */
    /**
     * Sample code: Create application security group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createApplicationSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationSecurityGroups()
            .createOrUpdate(
                "rg1", "test-asg", new ApplicationSecurityGroupInner().withLocation("westus"), Context.NONE);
    }
}
```
