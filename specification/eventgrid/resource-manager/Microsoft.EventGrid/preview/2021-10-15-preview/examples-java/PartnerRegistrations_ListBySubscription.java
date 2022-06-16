import com.azure.core.util.Context;

/** Samples for PartnerRegistrations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerRegistrations_ListBySubscription.json
     */
    /**
     * Sample code: PartnerRegistrations_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsListBySubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerRegistrations().list(null, null, Context.NONE);
    }
}
