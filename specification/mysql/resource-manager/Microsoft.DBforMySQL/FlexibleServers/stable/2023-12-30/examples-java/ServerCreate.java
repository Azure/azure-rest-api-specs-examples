
import com.azure.resourcemanager.mysqlflexibleserver.models.Backup;
import com.azure.resourcemanager.mysqlflexibleserver.models.CreateMode;
import com.azure.resourcemanager.mysqlflexibleserver.models.EnableStatusEnum;
import com.azure.resourcemanager.mysqlflexibleserver.models.HighAvailability;
import com.azure.resourcemanager.mysqlflexibleserver.models.HighAvailabilityMode;
import com.azure.resourcemanager.mysqlflexibleserver.models.MySqlServerSku;
import com.azure.resourcemanager.mysqlflexibleserver.models.ServerSkuTier;
import com.azure.resourcemanager.mysqlflexibleserver.models.ServerVersion;
import com.azure.resourcemanager.mysqlflexibleserver.models.Storage;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/ServerCreate
     * .json
     */
    /**
     * Sample code: Create a new server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void createANewServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().define("mysqltestserver").withRegion("southeastasia").withExistingResourceGroup("testrg")
            .withTags(mapOf("num", "1"))
            .withSku(new MySqlServerSku().withName("Standard_D2ds_v4").withTier(ServerSkuTier.GENERAL_PURPOSE))
            .withAdministratorLogin("cloudsa").withAdministratorLoginPassword("your_password")
            .withVersion(ServerVersion.FIVE_SEVEN).withAvailabilityZone("1").withCreateMode(CreateMode.DEFAULT)
            .withStorage(new Storage().withStorageSizeGB(100).withIops(600).withAutoGrow(EnableStatusEnum.DISABLED))
            .withBackup(new Backup().withBackupRetentionDays(7).withBackupIntervalHours(24)
                .withGeoRedundantBackup(EnableStatusEnum.DISABLED))
            .withHighAvailability(
                new HighAvailability().withMode(HighAvailabilityMode.ZONE_REDUNDANT).withStandbyAvailabilityZone("3"))
            .create();
    }

    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/stable/2023-12-30/examples/
     * ServerCreateWithPointInTimeRestore.json
     */
    /**
     * Sample code: Create a server as a point in time restore.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void
        createAServerAsAPointInTimeRestore(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().define("targetserver").withRegion("SoutheastAsia")
            .withExistingResourceGroup("TargetResourceGroup").withTags(mapOf("num", "1"))
            .withSku(new MySqlServerSku().withName("Standard_D14_v2").withTier(ServerSkuTier.GENERAL_PURPOSE))
            .withCreateMode(CreateMode.POINT_IN_TIME_RESTORE)
            .withSourceServerResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/sourceserver")
            .withRestorePointInTime(OffsetDateTime.parse("2021-06-24T00:00:37.467Z")).create();
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
