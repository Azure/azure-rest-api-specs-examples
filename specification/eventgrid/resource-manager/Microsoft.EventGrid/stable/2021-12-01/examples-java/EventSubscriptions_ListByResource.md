Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.4/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_ListByResource.json
     */
    /**
     * Sample code: EventSubscriptions_ListByResource.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListByResource(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .listByResource("examplerg", "Microsoft.EventGrid", "topics", "exampletopic2", null, null, Context.NONE);
    }
}
```
