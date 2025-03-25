
/**
 * Samples for PartnerTopics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PartnerTopics_Delete.json
     */
    /**
     * Sample code: PartnerTopics_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().delete("examplerg", "examplePartnerTopicName1", com.azure.core.util.Context.NONE);
    }
}
