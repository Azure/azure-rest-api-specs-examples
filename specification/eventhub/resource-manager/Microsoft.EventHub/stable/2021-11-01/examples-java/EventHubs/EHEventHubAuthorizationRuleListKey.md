Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EventHubs ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubAuthorizationRuleListKey.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleListKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubAuthorizationRuleListKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getEventHubs()
            .listKeysWithResponse(
                "ArunMonocle", "sdk-namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", Context.NONE);
    }
}
```
