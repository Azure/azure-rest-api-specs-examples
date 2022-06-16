import com.azure.core.util.Context;

/** Samples for EventChannels Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/EventChannels_Get.json
     */
    /**
     * Sample code: EventChannels_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventChannelsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventChannels()
            .getWithResponse("examplerg", "examplePartnerNamespaceName1", "exampleEventChannelName1", Context.NONE);
    }
}
