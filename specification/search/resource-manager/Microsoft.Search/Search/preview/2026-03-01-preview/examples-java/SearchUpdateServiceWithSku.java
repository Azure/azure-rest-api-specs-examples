
import com.azure.resourcemanager.search.models.SearchServiceUpdate;
import com.azure.resourcemanager.search.models.Sku;
import com.azure.resourcemanager.search.models.SkuName;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchUpdateServiceWithSku.json
     */
    /**
     * Sample code: SearchUpdateServiceWithSku.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchUpdateServiceWithSku(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().updateWithResponse("rg1", "mysearchservice",
            new SearchServiceUpdate().withSku(new Sku().withName(SkuName.STANDARD2))
                .withTags(mapOf("app-name", "My e-commerce app", "new-tag", "Adding a new tag")),
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
