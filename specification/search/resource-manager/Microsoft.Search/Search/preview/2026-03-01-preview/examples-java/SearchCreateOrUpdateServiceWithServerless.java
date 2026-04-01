
import com.azure.resourcemanager.search.fluent.models.SearchServiceInner;
import com.azure.resourcemanager.search.models.HostingMode;
import com.azure.resourcemanager.search.models.Sku;
import com.azure.resourcemanager.search.models.SkuName;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchCreateOrUpdateServiceWithServerless.json
     */
    /**
     * Sample code: SearchCreateOrUpdateServiceWithServerless.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        searchCreateOrUpdateServiceWithServerless(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().createOrUpdate("rg1", "myserverlessservice",
            new SearchServiceInner().withLocation("westus").withTags(mapOf("app-name", "My e-commerce app"))
                .withSku(new Sku().withName(SkuName.SERVERLESS)).withHostingMode(HostingMode.DEFAULT),
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
