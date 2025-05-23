
import com.azure.resourcemanager.datafactory.models.Factory;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Factories UpdateSync.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_Update.json
     */
    /**
     * Sample code: Factories_Update.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        Factory resource = manager.factories().getByResourceGroupWithResponse("exampleResourceGroup",
            "exampleFactoryName", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("exampleTag", "exampleValue")).apply();
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
