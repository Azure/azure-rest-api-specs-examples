import com.azure.core.util.Context;

/** Samples for PartnerRegistrations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/PartnerRegistrations_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerRegistrations_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsListByResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerRegistrations().listByResourceGroup("examplerg", null, null, Context.NONE);
    }
}
