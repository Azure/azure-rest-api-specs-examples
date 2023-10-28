import com.azure.resourcemanager.search.models.Identity;
import com.azure.resourcemanager.search.models.IdentityType;
import com.azure.resourcemanager.search.models.SearchServiceUpdate;
import com.azure.resourcemanager.search.models.Sku;
import com.azure.resourcemanager.search.models.SkuName;
import java.util.HashMap;
import java.util.Map;

/** Samples for Services Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchUpdateServiceToRemoveIdentity.json
     */
    /**
     * Sample code: SearchUpdateServiceToRemoveIdentity.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchUpdateServiceToRemoveIdentity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .searchServices()
            .manager()
            .serviceClient()
            .getServices()
            .updateWithResponse(
                "rg1",
                "mysearchservice",
                new SearchServiceUpdate()
                    .withSku(new Sku().withName(SkuName.STANDARD))
                    .withIdentity(new Identity().withType(IdentityType.NONE)),
                null,
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
