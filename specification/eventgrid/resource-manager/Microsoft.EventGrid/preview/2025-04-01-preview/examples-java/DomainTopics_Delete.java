
/**
 * Samples for DomainTopics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * DomainTopics_Delete.json
     */
    /**
     * Sample code: DomainTopics_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().delete("examplerg", "exampledomain1", "exampledomaintopic1",
            com.azure.core.util.Context.NONE);
    }
}
