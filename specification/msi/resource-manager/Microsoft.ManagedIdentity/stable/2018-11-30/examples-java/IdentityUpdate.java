import com.azure.core.util.Context;
import com.azure.resourcemanager.msi.models.IdentityUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for UserAssignedIdentities Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2018-11-30/examples/IdentityUpdate.json
     */
    /**
     * Sample code: IdentityUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void identityUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .identities()
            .manager()
            .serviceClient()
            .getUserAssignedIdentities()
            .updateWithResponse(
                "rgName",
                "resourceName",
                new IdentityUpdate().withLocation("eastus").withTags(mapOf("key1", "value1", "key2", "value2")),
                Context.NONE);
    }

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
