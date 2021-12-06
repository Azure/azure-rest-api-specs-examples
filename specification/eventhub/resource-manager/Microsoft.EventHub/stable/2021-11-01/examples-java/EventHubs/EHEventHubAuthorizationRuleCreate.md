Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.fluent.models.AuthorizationRuleInner;
import com.azure.resourcemanager.eventhubs.models.AccessRights;
import java.util.Arrays;

/** Samples for EventHubs CreateOrUpdateAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubAuthorizationRuleCreate.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubAuthorizationRuleCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getEventHubs()
            .createOrUpdateAuthorizationRuleWithResponse(
                "ArunMonocle",
                "sdk-Namespace-960",
                "sdk-EventHub-532",
                "sdk-Authrules-2513",
                new AuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
                Context.NONE);
    }
}
```
