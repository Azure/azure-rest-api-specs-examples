/** Samples for PartnerRegistrations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerRegistrations_ListBySubscription.json
     */
    /**
     * Sample code: PartnerRegistrations_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsListBySubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerRegistrations().list(null, null, com.azure.core.util.Context.NONE);
    }
}
