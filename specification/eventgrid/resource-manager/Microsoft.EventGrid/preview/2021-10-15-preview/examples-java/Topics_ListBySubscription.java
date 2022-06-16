import com.azure.core.util.Context;

/** Samples for Topics List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/Topics_ListBySubscription.json
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
