
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ServiceFabrics CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/
     * ServiceFabrics_CreateOrUpdate.json
     */
    /**
     * Sample code: ServiceFabrics_CreateOrUpdate.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricsCreateOrUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.serviceFabrics().define("{serviceFabricName}").withRegion("{location}")
            .withExistingUser("resourceGroupName", "{labName}", "{userName}").withTags(mapOf("tagName1", "tagValue1"))
            .withExternalServiceFabricId("{serviceFabricId}").withEnvironmentId("{environmentId}").create();
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
