import com.azure.core.util.Context;

/** Samples for PartnerTopics Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerTopics_Delete.json
     */
    /**
     * Sample code: PartnerTopics_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().delete("examplerg", "examplePartnerTopicName1", Context.NONE);
    }
}
