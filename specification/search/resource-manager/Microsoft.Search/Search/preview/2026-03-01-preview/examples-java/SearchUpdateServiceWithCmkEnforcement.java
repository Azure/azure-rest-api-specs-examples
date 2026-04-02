
import com.azure.resourcemanager.search.models.EncryptionWithCmk;
import com.azure.resourcemanager.search.models.SearchEncryptionWithCmk;
import com.azure.resourcemanager.search.models.SearchServiceUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchUpdateServiceWithCmkEnforcement.json
     */
    /**
     * Sample code: SearchUpdateServiceWithCmkEnforcement.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        searchUpdateServiceWithCmkEnforcement(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().updateWithResponse("rg1", "mysearchservice",
            new SearchServiceUpdate().withTags(mapOf("app-name", "My e-commerce app", "new-tag", "Adding a new tag"))
                .withReplicaCount(2)
                .withEncryptionWithCmk(new EncryptionWithCmk().withEnforcement(SearchEncryptionWithCmk.ENABLED)),
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
