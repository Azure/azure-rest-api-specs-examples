Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SystemTopics ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/SystemTopics_ListByResourceGroup.json
     */
    /**
     * Sample code: SystemTopics_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().listByResourceGroup("examplerg", null, null, Context.NONE);
    }
}
```
