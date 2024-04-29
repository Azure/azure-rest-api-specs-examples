
/**
 * Samples for DomainEventSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * DomainEventSubscriptions_Delete.json
     */
    /**
     * Sample code: DomainEventSubscriptions_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainEventSubscriptions().delete("examplerg", "exampleDomain1", "examplesubscription1",
            com.azure.core.util.Context.NONE);
    }
}
