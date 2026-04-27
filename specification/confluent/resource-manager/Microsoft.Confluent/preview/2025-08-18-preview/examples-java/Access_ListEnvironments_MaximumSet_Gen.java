
import com.azure.resourcemanager.confluent.models.ListAccessRequestModel;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Access ListEnvironments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Access_ListEnvironments_MaximumSet_Gen.json
     */
    /**
     * Sample code: Access_ListEnvironments_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void accessListEnvironmentsMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.access().listEnvironmentsWithResponse("rgconfluent", "rnbjtcdqddweb",
            new ListAccessRequestModel().withSearchFilters(mapOf("key8083", "fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
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
