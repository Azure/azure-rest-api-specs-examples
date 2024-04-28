
import com.azure.resourcemanager.eventgrid.models.ChannelUpdateParameters;
import java.time.OffsetDateTime;

/**
 * Samples for Channels Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/Channels_Update.
     * json
     */
    /**
     * Sample code: Channels_Update.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.channels().updateWithResponse(
            "examplerg", "examplePartnerNamespaceName1", "exampleChannelName1", new ChannelUpdateParameters()
                .withExpirationTimeIfNotActivatedUtc(OffsetDateTime.parse("2022-03-23T23:06:11.785Z")),
            com.azure.core.util.Context.NONE);
    }
}
