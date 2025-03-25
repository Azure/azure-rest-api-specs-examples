
/**
 * Samples for PartnerTopics Deactivate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PartnerTopics_Deactivate.
     * json
     */
    /**
     * Sample code: PartnerTopics_Deactivate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsDeactivate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().deactivateWithResponse("examplerg", "examplePartnerTopic1",
            com.azure.core.util.Context.NONE);
    }
}
