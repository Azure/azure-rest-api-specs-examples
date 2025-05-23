
/**
 * Samples for PartnerTopics ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerTopics_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerTopics_ListByResourceGroup.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().listByResourceGroup("examplerg", null, null, com.azure.core.util.Context.NONE);
    }
}
