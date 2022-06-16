import com.azure.core.util.Context;

/** Samples for DomainTopicEventSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainTopicEventSubscriptions_Get.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicEventSubscriptionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domainTopicEventSubscriptions()
            .getWithResponse(
                "examplerg", "exampleDomain1", "exampleDomainTopic1", "examplesubscription1", Context.NONE);
    }
}
