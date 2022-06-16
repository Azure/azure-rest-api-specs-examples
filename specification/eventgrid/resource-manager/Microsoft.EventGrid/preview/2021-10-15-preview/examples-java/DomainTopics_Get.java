import com.azure.core.util.Context;

/** Samples for DomainTopics Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainTopics_Get.json
     */
    /**
     * Sample code: DomainTopics_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().getWithResponse("examplerg", "exampledomain2", "topic1", Context.NONE);
    }
}
