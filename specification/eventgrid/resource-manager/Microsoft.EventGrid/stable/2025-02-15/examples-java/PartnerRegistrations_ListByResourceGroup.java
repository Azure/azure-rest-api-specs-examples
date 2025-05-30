
/**
 * Samples for PartnerRegistrations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerRegistrations_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerRegistrations_ListByResourceGroup.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        partnerRegistrationsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerRegistrations().listByResourceGroup("examplerg", null, null, com.azure.core.util.Context.NONE);
    }
}
