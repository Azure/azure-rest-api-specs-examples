Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.4/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SystemTopics Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/SystemTopics_Delete.json
     */
    /**
     * Sample code: SystemTopics_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().delete("examplerg", "exampleSystemTopic1", Context.NONE);
    }
}
```
