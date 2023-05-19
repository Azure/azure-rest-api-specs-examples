/** Samples for PartnerDestinations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerDestinations_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerDestinations_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerDestinationsListByResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerDestinations().listByResourceGroup("examplerg", null, null, com.azure.core.util.Context.NONE);
    }
}
