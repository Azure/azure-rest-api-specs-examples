
/**
 * Samples for Topics ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/
     * SBTopicListByNameSpace.json
     */
    /**
     * Sample code: TopicGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void topicGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getTopics().listByNamespace("Default-ServiceBus-WestUS",
            "sdk-Namespace-1617", null, null, com.azure.core.util.Context.NONE);
    }
}
