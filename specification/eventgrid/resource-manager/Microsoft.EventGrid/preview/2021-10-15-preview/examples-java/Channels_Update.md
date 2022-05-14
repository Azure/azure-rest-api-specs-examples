Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.ChannelUpdateParameters;
import java.time.OffsetDateTime;

/** Samples for Channels Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/Channels_Update.json
     */
    /**
     * Sample code: Channels_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .channels()
            .updateWithResponse(
                "examplerg",
                "examplePartnerNamespaceName1",
                "exampleChannelName1",
                new ChannelUpdateParameters()
                    .withExpirationTimeIfNotActivatedUtc(OffsetDateTime.parse("2022-03-23T23:06:11.785Z")),
                Context.NONE);
    }
}
```
