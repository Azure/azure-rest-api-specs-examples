import com.azure.core.util.Context;

/** Samples for TopicTypes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/TopicTypes_List.json
     */
    /**
     * Sample code: TopicTypes_List.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicTypesList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicTypes().list(Context.NONE);
    }
}
