import com.azure.core.util.Context;

/** Samples for SystemTopics ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopics_ListByResourceGroup.json
     */
    /**
     * Sample code: SystemTopics_ListByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsListByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().listByResourceGroup("examplerg", null, null, Context.NONE);
    }
}
