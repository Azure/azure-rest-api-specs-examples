/** Samples for PartnerTopics Activate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerTopics_Activate.json
     */
    /**
     * Sample code: PartnerTopics_Activate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsActivate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerTopics()
            .activateWithResponse("examplerg", "examplePartnerTopic1", com.azure.core.util.Context.NONE);
    }
}
