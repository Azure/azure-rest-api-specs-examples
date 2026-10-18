
import com.azure.resourcemanager.storage.fluent.models.ContextCacheInner;
import com.azure.resourcemanager.storage.models.ContextCacheAccountKind;
import com.azure.resourcemanager.storage.models.ContextCacheProperties;
import com.azure.resourcemanager.storage.models.CustomerManagedKeyEncryption;
import com.azure.resourcemanager.storage.models.Identity;
import com.azure.resourcemanager.storage.models.IdentityType;
import com.azure.resourcemanager.storage.models.KeyEncryptionKeyIdentity;
import com.azure.resourcemanager.storage.models.KeyEncryptionKeyIdentityType;
import com.azure.resourcemanager.storage.models.StorageAccountEncryption;
import com.azure.resourcemanager.storage.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ContextCaches CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/StorageContextCacheCRUD/ContextCaches_CreateOrUpdate_UserAssignedIdentity.json
     */
    /**
     * Sample code: Create a Context Cache with user-assigned identity.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        createAContextCacheWithUserAssignedIdentity(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCaches().createOrUpdate("testrg", "testcontextcache", new ContextCacheInner()
            .withLocation("eastus").withTags(mapOf("environment", "test"))
            .withProperties(new ContextCacheProperties().withAccountKind(ContextCacheAccountKind.REGIONAL)
                .withDescription("Test context cache")
                .withEncryption(new StorageAccountEncryption().withCustomerManagedKeyEncryption(
                    new CustomerManagedKeyEncryption().withKeyEncryptionKeyIdentity(new KeyEncryptionKeyIdentity()
                        .withIdentityType(KeyEncryptionKeyIdentityType.USER_ASSIGNED_IDENTITY)
                        .withUserAssignedIdentityResourceId(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-uami"))
                        .withKeyEncryptionKeyUrl("fakeTokenPlaceholder"))))
            .withIdentity(new Identity().withType(IdentityType.USER_ASSIGNED).withUserAssignedIdentities(mapOf(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-uami",
                new UserAssignedIdentity()))),
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
