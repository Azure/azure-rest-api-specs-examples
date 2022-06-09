```java
import com.azure.core.util.Context;

/** Samples for EventChannels Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/EventChannels_Delete.json
     */
    /**
     * Sample code: EventChannels_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventChannelsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventChannels()
            .delete("examplerg", "examplePartnerNamespaceName1", "exampleEventChannelName", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.
