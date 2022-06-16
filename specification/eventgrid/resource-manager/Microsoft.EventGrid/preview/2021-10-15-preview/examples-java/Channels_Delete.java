import com.azure.core.util.Context;

/** Samples for Channels Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/Channels_Delete.json
     */
    /**
     * Sample code: Channels_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .channels()
            .delete("examplerg", "examplePartnerNamespaceName1", "exampleEventChannelName1", Context.NONE);
    }
}
