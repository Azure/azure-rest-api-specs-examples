
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.DatabaseIdentity;
import com.azure.resourcemanager.sql.models.DatabaseIdentityType;
import com.azure.resourcemanager.sql.models.DatabaseUserIdentity;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/UpdateDatabaseHyperscaleMigrationWithManualCutover.json
     */
    /**
     * Sample code: Updates a database to Hyperscale SLO with manual cutover.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updatesADatabaseToHyperscaleSLOWithManualCutover(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new DatabaseInner().withLocation("southeastasia")
                .withSku(new Sku().withName("HS_Gen5_2").withTier("Hyperscale"))
                .withIdentity(new DatabaseIdentity().withType(DatabaseIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(mapOf(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/userAssignedIdentities/umi",
                        new DatabaseUserIdentity(),
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/userAssignedIdentities/umiToDelete",
                        null)))
                .withManualCutover(true),
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
