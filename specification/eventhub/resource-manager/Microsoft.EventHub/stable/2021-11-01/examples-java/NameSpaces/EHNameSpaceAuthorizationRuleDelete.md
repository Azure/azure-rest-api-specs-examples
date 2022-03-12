Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Namespaces DeleteAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceAuthorizationRuleDelete.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceAuthorizationRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getNamespaces()
            .deleteAuthorizationRuleWithResponse(
                "ArunMonocle", "sdk-Namespace-8980", "sdk-Authrules-8929", Context.NONE);
    }
}
```
