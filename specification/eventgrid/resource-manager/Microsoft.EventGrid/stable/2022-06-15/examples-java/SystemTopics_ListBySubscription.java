import com.azure.core.util.Context;

/** Samples for SystemTopics List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopics_ListBySubscription.json
     */
    /**
     * Sample code: SystemTopics_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsListBySubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().list(null, null, Context.NONE);
    }
}
