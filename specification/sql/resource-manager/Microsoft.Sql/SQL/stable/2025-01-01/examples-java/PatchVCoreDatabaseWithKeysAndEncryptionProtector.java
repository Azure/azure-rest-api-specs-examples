
import com.azure.resourcemanager.sql.models.DatabaseIdentity;
import com.azure.resourcemanager.sql.models.DatabaseIdentityType;
import com.azure.resourcemanager.sql.models.DatabaseKey;
import com.azure.resourcemanager.sql.models.DatabaseUpdate;
import com.azure.resourcemanager.sql.models.DatabaseUserIdentity;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/PatchVCoreDatabaseWithKeysAndEncryptionProtector.json
     */
    /**
     * Sample code: Patch a database with database-level customer managed keys.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        patchADatabaseWithDatabaseLevelCustomerManagedKeys(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().update("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new DatabaseUpdate().withSku(new Sku().withName("S0").withTier("Standard")).withIdentity(
                new DatabaseIdentity().withType(DatabaseIdentityType.USER_ASSIGNED).withUserAssignedIdentities(mapOf(
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi",
                    new DatabaseUserIdentity(),
                    "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umiToDelete",
                    null)))
                .withKeys(mapOf("https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion", new DatabaseKey(),
                    "https://your-key-vault-name.vault.azure.net/yourKey2/yourKey2VersionToDelete", null))
                .withEncryptionProtector("https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion"),
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
