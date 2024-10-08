
import com.azure.resourcemanager.msi.fluent.models.IdentityInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for UserAssignedIdentities CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/IdentityCreate.json
     */
    /**
     * Sample code: IdentityCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void identityCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getUserAssignedIdentities().createOrUpdateWithResponse("rgName",
            "resourceName",
            new IdentityInner().withLocation("eastus")
                .withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder")),
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
