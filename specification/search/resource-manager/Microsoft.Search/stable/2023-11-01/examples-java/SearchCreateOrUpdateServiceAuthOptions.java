
import com.azure.resourcemanager.search.fluent.models.SearchServiceInner;
import com.azure.resourcemanager.search.models.AadAuthFailureMode;
import com.azure.resourcemanager.search.models.DataPlaneAadOrApiKeyAuthOption;
import com.azure.resourcemanager.search.models.DataPlaneAuthOptions;
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
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/
     * SearchCreateOrUpdateServiceAuthOptions.json
     */
    /**
     * Sample code: SearchCreateOrUpdateServiceAuthOptions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchCreateOrUpdateServiceAuthOptions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getServices().createOrUpdate("rg1", "mysearchservice",
            new SearchServiceInner().withLocation("westus").withTags(mapOf("app-name", "My e-commerce app"))
                .withSku(new Sku().withName(SkuName.STANDARD)).withReplicaCount(3).withPartitionCount(1)
                .withHostingMode(HostingMode.DEFAULT)
                .withAuthOptions(new DataPlaneAuthOptions().withAadOrApiKey(new DataPlaneAadOrApiKeyAuthOption()
                    .withAadAuthFailureMode(AadAuthFailureMode.HTTP401WITH_BEARER_CHALLENGE))),
            null, com.azure.core.util.Context.NONE);
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
