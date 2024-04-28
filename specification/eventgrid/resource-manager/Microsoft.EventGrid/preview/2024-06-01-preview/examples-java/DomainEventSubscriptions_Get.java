
/**
 * Samples for DomainEventSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * DomainEventSubscriptions_Get.json
     */
    /**
     * Sample code: DomainEventSubscriptions_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainEventSubscriptions().getWithResponse("examplerg", "exampleDomain1", "examplesubscription1",
            com.azure.core.util.Context.NONE);
    }
}
