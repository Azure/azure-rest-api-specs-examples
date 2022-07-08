import com.azure.core.util.Context;

/** Samples for DomainEventSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainEventSubscriptions_List.json
     */
    /**
     * Sample code: DomainEventSubscriptions_List.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainEventSubscriptions().list("examplerg", "exampleDomain1", null, null, Context.NONE);
    }
}
