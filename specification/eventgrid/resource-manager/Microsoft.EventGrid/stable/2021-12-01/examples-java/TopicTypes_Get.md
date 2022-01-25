Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for TopicTypes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/TopicTypes_Get.json
     */
    /**
     * Sample code: TopicTypes_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicTypesGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicTypes().getWithResponse("Microsoft.Storage.StorageAccounts", Context.NONE);
    }
}
```
