Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for NatRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NatRuleDelete.json
     */
    /**
     * Sample code: NatRuleDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void natRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNatRules().delete("rg1", "gateway1", "natRule1", Context.NONE);
    }
}
```
