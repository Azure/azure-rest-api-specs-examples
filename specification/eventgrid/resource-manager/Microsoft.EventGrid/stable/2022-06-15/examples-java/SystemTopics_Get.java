import com.azure.core.util.Context;

/** Samples for SystemTopics GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopics_Get.json
     */
    /**
     * Sample code: SystemTopics_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().getByResourceGroupWithResponse("examplerg", "exampleSystemTopic2", Context.NONE);
    }
}
