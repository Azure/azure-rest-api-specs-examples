import com.azure.core.util.Context;

/** Samples for DomainTopics ListByDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainTopics_ListByDomain.json
     */
    /**
     * Sample code: DomainTopics_ListByDomain.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsListByDomain(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().listByDomain("examplerg", "exampledomain2", null, null, Context.NONE);
    }
}
