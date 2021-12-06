Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.4/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SystemTopicEventSubscriptions ListBySystemTopic. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/SystemTopicEventSubscriptions_ListBySystemTopic.json
     */
    /**
     * Sample code: SystemTopicEventSubscriptions_ListBySystemTopic.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicEventSubscriptionsListBySystemTopic(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .systemTopicEventSubscriptions()
            .listBySystemTopic("examplerg", "exampleSystemTopic1", null, null, Context.NONE);
    }
}
```
