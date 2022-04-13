Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DefaultSecurityRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DefaultSecurityRuleGet.json
     */
    /**
     * Sample code: DefaultSecurityRuleGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void defaultSecurityRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getDefaultSecurityRules()
            .getWithResponse("testrg", "nsg1", "AllowVnetInBound", Context.NONE);
    }
}
```
