import com.azure.core.util.Context;

/** Samples for PartnerTopics Activate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerTopics_Activate.json
     */
    /**
     * Sample code: PartnerTopics_Activate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsActivate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().activateWithResponse("examplerg", "examplePartnerTopic1", Context.NONE);
    }
}
