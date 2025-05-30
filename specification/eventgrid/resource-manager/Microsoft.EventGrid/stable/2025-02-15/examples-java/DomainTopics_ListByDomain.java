
/**
 * Samples for DomainTopics ListByDomain.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/DomainTopics_ListByDomain
     * .json
     */
    /**
     * Sample code: DomainTopics_ListByDomain.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsListByDomain(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().listByDomain("examplerg", "exampledomain2", null, null,
            com.azure.core.util.Context.NONE);
    }
}
