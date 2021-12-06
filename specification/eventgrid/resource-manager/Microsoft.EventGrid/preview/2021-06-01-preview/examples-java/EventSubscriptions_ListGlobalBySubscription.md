Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EventSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/EventSubscriptions_ListGlobalBySubscription.json
     */
    /**
     * Sample code: EventSubscriptions_ListGlobalBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListGlobalBySubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().list(null, null, Context.NONE);
    }
}
```
