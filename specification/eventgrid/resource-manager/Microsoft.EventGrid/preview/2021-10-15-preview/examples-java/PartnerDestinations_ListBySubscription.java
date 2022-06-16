import com.azure.core.util.Context;

/** Samples for PartnerDestinations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerDestinations_ListBySubscription.json
     */
    /**
     * Sample code: PartnerDestinations_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerDestinationsListBySubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerDestinations().list(null, null, Context.NONE);
    }
}
