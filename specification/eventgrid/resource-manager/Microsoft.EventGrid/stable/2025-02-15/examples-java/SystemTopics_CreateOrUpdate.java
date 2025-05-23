
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SystemTopics CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * SystemTopics_CreateOrUpdate.json
     */
    /**
     * Sample code: SystemTopics_CreateOrUpdate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopics().define("exampleSystemTopic1").withRegion("westus2")
            .withExistingResourceGroup("examplerg").withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withSource(
                "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/microsoft.storage/storageaccounts/pubstgrunnerb71cd29e")
            .withTopicType("microsoft.storage.storageaccounts").create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
