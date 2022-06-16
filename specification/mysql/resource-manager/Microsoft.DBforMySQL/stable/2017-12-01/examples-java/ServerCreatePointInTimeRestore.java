import com.azure.resourcemanager.mysql.models.ServerPropertiesForRestore;
import com.azure.resourcemanager.mysql.models.Sku;
import com.azure.resourcemanager.mysql.models.SkuTier;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerCreatePointInTimeRestore.json
     */
    /**
     * Sample code: Create a database as a point in time restore.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createADatabaseAsAPointInTimeRestore(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .servers()
            .define("targetserver")
            .withRegion("brazilsouth")
            .withExistingResourceGroup("TargetResourceGroup")
            .withProperties(
                new ServerPropertiesForRestore()
                    .withSourceServerId(
                        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/servers/sourceserver")
                    .withRestorePointInTime(OffsetDateTime.parse("2017-12-14T00:00:37.467Z")))
            .withTags(mapOf("ElasticServer", "1"))
            .withSku(
                new Sku().withName("GP_Gen5_2").withTier(SkuTier.GENERAL_PURPOSE).withCapacity(2).withFamily("Gen5"))
            .create();
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
