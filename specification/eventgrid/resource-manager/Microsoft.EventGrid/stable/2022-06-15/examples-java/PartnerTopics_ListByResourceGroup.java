import com.azure.core.util.Context;

/** Samples for PartnerTopics ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerTopics_ListByResourceGroup.json
     */
    /**
     * Sample code: PartnerTopics_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().listByResourceGroup("examplerg", null, null, Context.NONE);
    }
}
