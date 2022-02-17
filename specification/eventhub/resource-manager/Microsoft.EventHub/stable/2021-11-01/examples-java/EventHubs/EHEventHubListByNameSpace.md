Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EventHubs ListByNamespace. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubListByNameSpace.json
     */
    /**
     * Sample code: EventHubsListAll.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubsListAll(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getEventHubs()
            .listByNamespace("Default-NotificationHubs-AustraliaEast", "sdk-Namespace-5357", null, null, Context.NONE);
    }
}
```
