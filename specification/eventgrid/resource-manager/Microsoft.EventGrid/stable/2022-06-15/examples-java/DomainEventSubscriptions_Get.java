import com.azure.core.util.Context;

/** Samples for DomainEventSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainEventSubscriptions_Get.json
     */
    /**
     * Sample code: DomainEventSubscriptions_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domainEventSubscriptions()
            .getWithResponse("examplerg", "exampleDomain1", "examplesubscription1", Context.NONE);
    }
}
