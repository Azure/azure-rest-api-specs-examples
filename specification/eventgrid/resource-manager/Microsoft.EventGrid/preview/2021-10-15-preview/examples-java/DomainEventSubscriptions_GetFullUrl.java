import com.azure.core.util.Context;

/** Samples for DomainEventSubscriptions GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainEventSubscriptions_GetFullUrl.json
     */
    /**
     * Sample code: DomainEventSubscriptions_GetFullUrl.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsGetFullUrl(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domainEventSubscriptions()
            .getFullUrlWithResponse("examplerg", "exampleDomain1", "examplesubscription1", Context.NONE);
    }
}
