import com.azure.core.util.Context;

/** Samples for Topics Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/Topics_Delete.json
     */
    /**
     * Sample code: Topics_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topics().delete("examplerg", "exampletopic1", Context.NONE);
    }
}
