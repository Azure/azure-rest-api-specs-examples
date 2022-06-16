import com.azure.core.util.Context;

/** Samples for EventChannels ListByPartnerNamespace. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventChannels_ListByPartnerNamespace.json
     */
    /**
     * Sample code: EventChannels_ListByPartnerNamespace.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventChannelsListByPartnerNamespace(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventChannels().listByPartnerNamespace("examplerg", "partnerNamespace123", null, null, Context.NONE);
    }
}
