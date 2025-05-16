
import com.azure.resourcemanager.postgresqlflexibleserver.models.AzureManagedDiskPerformanceTiers;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Backup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.GeoRedundantBackupEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.HighAvailability;
import com.azure.resourcemanager.postgresqlflexibleserver.models.HighAvailabilityMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Network;
import com.azure.resourcemanager.postgresqlflexibleserver.models.ServerVersion;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Sku;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Storage;
import com.azure.resourcemanager.postgresqlflexibleserver.models.StorageAutoGrow;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * ServerCreate.json
     */
    /**
     * Sample code: Create a new server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void createANewServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().define("testpgflex").withRegion("eastus").withExistingResourceGroup("testrg")
            .withTags(mapOf("VNetServer", "1"))
            .withSku(new Sku().withName("Standard_D4ds_v5").withTier(SkuTier.GENERAL_PURPOSE))
            .withAdministratorLogin("login").withAdministratorLoginPassword("Password1")
            .withVersion(ServerVersion.ONE_SIX)
            .withStorage(new Storage().withStorageSizeGB(512).withAutoGrow(StorageAutoGrow.DISABLED)
                .withTier(AzureManagedDiskPerformanceTiers.P20))
            .withBackup(new Backup().withBackupRetentionDays(7).withGeoRedundantBackup(GeoRedundantBackupEnum.ENABLED))
            .withNetwork(new Network().withDelegatedSubnetResourceId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/test-subnet")
                .withPrivateDnsZoneArmResourceId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.Network/privateDnsZones/testpgflex.private.postgres.database"))
            .withHighAvailability(new HighAvailability().withMode(HighAvailabilityMode.ZONE_REDUNDANT))
            .withAvailabilityZone("1").withCreateMode(CreateMode.CREATE).create();
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
