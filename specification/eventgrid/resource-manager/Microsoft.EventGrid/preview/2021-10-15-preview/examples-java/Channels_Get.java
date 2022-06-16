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
