
import com.azure.core.util.Context;
import com.azure.resourcemanager.servicebus.fluent.models.SBTopicInner;

/** Samples for Topics CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/SBTopicCreate.
     * json
     */
    /**
     * Sample code: TopicCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().createOrUpdateWithResponse("ArunMonocle",
            "sdk-Namespace-1617", "sdk-Topics-5488", new SBTopicInner().withEnableExpress(true), Context.NONE);
    }
}
