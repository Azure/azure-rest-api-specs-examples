Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SystemTopics List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/SystemTopics_ListBySubscription.json
     */
    /**
     * Sample code: SystemTopics_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsListBySubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().list(null, null, Context.NONE);
    }
}
```
