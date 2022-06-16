import com.azure.core.util.Context;

/** Samples for Topics List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/Topics_ListBySubscription.json
     */
    /**
     * Sample code: Topics_ListBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsListBySubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topics().list(null, null, Context.NONE);
    }
}
