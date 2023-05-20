import java.util.Arrays;

/** Samples for TopicSpaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/TopicSpaces_CreateOrUpdate.json
     */
    /**
     * Sample code: TopicSpaces_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicSpacesCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .topicSpaces()
            .define("exampleTopicSpaceName1")
            .withExistingNamespace("examplerg", "exampleNamespaceName1")
            .withTopicTemplates(Arrays.asList("filter1", "filter2"))
            .create();
    }
}
