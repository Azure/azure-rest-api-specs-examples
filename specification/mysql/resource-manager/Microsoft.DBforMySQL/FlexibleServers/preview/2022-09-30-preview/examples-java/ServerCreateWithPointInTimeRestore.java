import com.azure.resourcemanager.mysqlflexibleserver.models.CreateMode;
import com.azure.resourcemanager.mysqlflexibleserver.models.Sku;
import com.azure.resourcemanager.mysqlflexibleserver.models.SkuTier;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2022-09-30-preview/examples/ServerCreateWithPointInTimeRestore.json
     */
    /**
     * Sample code: Create a server as a point in time restore.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createAServerAsAPointInTimeRestore(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .servers()
            .define("targetserver")
            .withRegion("SoutheastAsia")
            .withExistingResourceGroup("TargetResourceGroup")
            .withTags(mapOf("num", "1"))
            .withSku(new Sku().withName("Standard_D14_v2").withTier(SkuTier.GENERAL_PURPOSE))
            .withCreateMode(CreateMode.POINT_IN_TIME_RESTORE)
            .withSourceServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/sourceserver")
            .withRestorePointInTime(OffsetDateTime.parse("2021-06-24T00:00:37.467Z"))
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
