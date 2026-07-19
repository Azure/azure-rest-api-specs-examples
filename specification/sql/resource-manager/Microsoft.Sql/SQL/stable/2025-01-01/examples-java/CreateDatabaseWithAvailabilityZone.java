
import com.azure.resourcemanager.sql.fluent.models.DatabaseInner;
import com.azure.resourcemanager.sql.models.AvailabilityZoneType;
import com.azure.resourcemanager.sql.models.CreateMode;
import com.azure.resourcemanager.sql.models.Sku;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Databases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateDatabaseWithAvailabilityZone.json
     */
    /**
     * Sample code: Creates a database with availability zone specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createsADatabaseWithAvailabilityZoneSpecified(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().createOrUpdate("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            new DatabaseInner().withLocation("southeastasia").withSku(new Sku().withName("S0").withTier("Standard"))
                .withCreateMode(CreateMode.DEFAULT).withCollation("SQL_Latin1_General_CP1_CI_AS")
                .withMaxSizeBytes(1073741824L).withAvailabilityZone(AvailabilityZoneType.ONE),
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
