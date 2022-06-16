import com.azure.core.util.Context;
import com.azure.resourcemanager.webpubsub.models.NameAvailabilityParameters;

/** Samples for WebPubSub CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_CheckNameAvailability.json
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
