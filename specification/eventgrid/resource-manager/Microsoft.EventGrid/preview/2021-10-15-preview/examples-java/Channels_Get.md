Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.1/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Channels Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/Channels_Get.json
     */
    /**
     * Sample code: Channels_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .channels()
            .getWithResponse("examplerg", "examplePartnerNamespaceName1", "exampleChannelName1", Context.NONE);
    }
}
```
