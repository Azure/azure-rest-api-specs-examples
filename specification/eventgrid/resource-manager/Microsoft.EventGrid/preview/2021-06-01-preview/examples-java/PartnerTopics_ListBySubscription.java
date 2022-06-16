import com.azure.core.util.Context;

/** Samples for PartnerTopics List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/PartnerTopics_ListBySubscription.json
     */
    /**
     * Sample code: PartnerTopics_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsListBySubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopics().list(null, null, Context.NONE);
    }
}
