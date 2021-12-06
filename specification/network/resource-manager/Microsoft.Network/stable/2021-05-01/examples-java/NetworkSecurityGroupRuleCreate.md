Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.SecurityRuleInner;
import com.azure.resourcemanager.network.models.SecurityRuleAccess;
import com.azure.resourcemanager.network.models.SecurityRuleDirection;
import com.azure.resourcemanager.network.models.SecurityRuleProtocol;

/** Samples for SecurityRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkSecurityGroupRuleCreate.json
     */
    /**
     * Sample code: Create security rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createSecurityRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSecurityRules()
            .createOrUpdate(
                "rg1",
                "testnsg",
                "rule1",
                new SecurityRuleInner()
                    .withProtocol(SecurityRuleProtocol.ASTERISK)
                    .withSourcePortRange("*")
                    .withDestinationPortRange("8080")
                    .withSourceAddressPrefix("10.0.0.0/8")
                    .withDestinationAddressPrefix("11.0.0.0/8")
                    .withAccess(SecurityRuleAccess.DENY)
                    .withPriority(100)
                    .withDirection(SecurityRuleDirection.OUTBOUND),
                Context.NONE);
    }
}
```
