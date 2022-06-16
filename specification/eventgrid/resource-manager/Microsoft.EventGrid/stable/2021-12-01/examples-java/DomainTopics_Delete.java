import com.azure.core.util.Context;

/** Samples for DomainTopics Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/DomainTopics_Delete.json
     */
    /**
     * Sample code: DomainTopics_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().delete("examplerg", "exampledomain1", "exampledomaintopic1", Context.NONE);
    }
}
