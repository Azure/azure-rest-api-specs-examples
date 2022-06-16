import com.azure.core.util.Context;

/** Samples for Topics GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/Topics_Get.json
     */
    /**
     * Sample code: Topics_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topics().getByResourceGroupWithResponse("examplerg", "exampletopic2", Context.NONE);
    }
}
