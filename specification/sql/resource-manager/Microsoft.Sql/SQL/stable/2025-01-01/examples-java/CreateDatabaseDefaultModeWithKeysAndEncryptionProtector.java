
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.CreateMode;
import com.azure.resourcemanager.sql.models.DatabaseIdentity;
import com.azure.resourcemanager.sql.models.DatabaseIdentityType;
import com.azure.resourcemanager.sql.models.DatabaseKey;
import com.azure.resourcemanager.sql.models.DatabaseUserIdentity;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateDatabaseDefaultModeWithKeysAndEncryptionProtector.json
     */
    /**
     * Sample code: Creates a database with database-level customer managed keys.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createsADatabaseWithDatabaseLevelCustomerManagedKeys(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new DatabaseInner().withLocation("southeastasia").withSku(new Sku().withName("S0").withTier("Standard"))
                .withIdentity(new DatabaseIdentity().withType(DatabaseIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(mapOf(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi",
                        new DatabaseUserIdentity())))
                .withCreateMode(CreateMode.DEFAULT).withCollation("SQL_Latin1_General_CP1_CI_AS")
                .withMaxSizeBytes(1073741824L)
                .withKeys(mapOf("https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion", new DatabaseKey(),
                    "https://your-key-vault-name.vault.azure.net/yourKey2/yourKey2Version", new DatabaseKey()))
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
