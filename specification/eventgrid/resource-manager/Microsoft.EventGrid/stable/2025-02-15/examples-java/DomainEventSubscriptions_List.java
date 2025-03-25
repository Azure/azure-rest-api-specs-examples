
/**
 * Samples for DomainEventSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * DomainEventSubscriptions_List.json
     */
    /**
     * Sample code: DomainEventSubscriptions_List.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainEventSubscriptions().list("examplerg", "exampleDomain1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
