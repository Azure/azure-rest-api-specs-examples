
import com.azure.core.util.Context;

/** Samples for Topics Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/SBTopicDelete.
     * json
     */
    /**
     * Sample code: TopicDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().deleteWithResponse("ArunMonocle",
            "sdk-Namespace-1617", "sdk-Topics-5488", Context.NONE);
    }
}
