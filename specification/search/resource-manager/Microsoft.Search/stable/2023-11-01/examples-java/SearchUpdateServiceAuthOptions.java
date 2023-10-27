import com.azure.resourcemanager.search.models.AadAuthFailureMode;
import com.azure.resourcemanager.search.models.DataPlaneAadOrApiKeyAuthOption;
import com.azure.resourcemanager.search.models.DataPlaneAuthOptions;
import com.azure.resourcemanager.search.models.SearchServiceUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for Services Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchUpdateServiceAuthOptions.json
     */
    /**
     * Sample code: SearchUpdateServiceAuthOptions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchUpdateServiceAuthOptions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .searchServices()
            .manager()
            .serviceClient()
            .getServices()
            .updateWithResponse(
                "rg1",
                "mysearchservice",
                new SearchServiceUpdate()
                    .withTags(mapOf("app-name", "My e-commerce app", "new-tag", "Adding a new tag"))
                    .withReplicaCount(2)
                    .withAuthOptions(
                        new DataPlaneAuthOptions()
                            .withAadOrApiKey(
                                new DataPlaneAadOrApiKeyAuthOption()
                                    .withAadAuthFailureMode(AadAuthFailureMode.HTTP401WITH_BEARER_CHALLENGE))),
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
