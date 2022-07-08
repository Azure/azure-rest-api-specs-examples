import com.azure.core.util.Context;

/** Samples for DomainEventSubscriptions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainEventSubscriptions_Delete.json
     */
    /**
     * Sample code: DomainEventSubscriptions_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainEventSubscriptions().delete("examplerg", "exampleDomain1", "examplesubscription1", Context.NONE);
    }
}
