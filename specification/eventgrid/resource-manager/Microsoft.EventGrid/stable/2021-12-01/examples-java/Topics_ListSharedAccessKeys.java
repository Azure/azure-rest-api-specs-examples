import com.azure.core.util.Context;

/** Samples for Topics ListSharedAccessKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/Topics_ListSharedAccessKeys.json
     */
    /**
     * Sample code: Topics_ListSharedAccessKeys.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsListSharedAccessKeys(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topics().listSharedAccessKeysWithResponse("examplerg", "exampletopic2", Context.NONE);
    }
}
