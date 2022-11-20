import com.azure.core.util.Context;
import com.azure.resourcemanager.webpubsub.models.NameAvailabilityParameters;

/** Samples for WebPubSub CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSub_CheckNameAvailability.json
     */
    /**
     * Sample code: WebPubSub_CheckNameAvailability.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCheckNameAvailability(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubs()
            .checkNameAvailabilityWithResponse(
                "eastus",
                new NameAvailabilityParameters()
                    .withType("Microsoft.SignalRService/WebPubSub")
                    .withName("myWebPubSubService"),
                Context.NONE);
    }
}
