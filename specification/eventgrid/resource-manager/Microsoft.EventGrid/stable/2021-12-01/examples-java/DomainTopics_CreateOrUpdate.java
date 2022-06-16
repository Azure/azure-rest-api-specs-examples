import com.azure.core.util.Context;

/** Samples for DomainTopics CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/DomainTopics_CreateOrUpdate.json
     */
    /**
     * Sample code: DomainTopics_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopics().createOrUpdate("examplerg", "exampledomain1", "exampledomaintopic1", Context.NONE);
    }
}
