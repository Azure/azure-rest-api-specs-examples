
import com.azure.resourcemanager.postgresqlflexibleserver.models.AzureManagedDiskPerformanceTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Backup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Cluster;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.GeographicallyRedundantBackup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.HighAvailability;
import com.azure.resourcemanager.postgresqlflexibleserver.models.HighAvailabilityMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Network;
import com.azure.resourcemanager.postgresqlflexibleserver.models.PostgresMajorVersion;
import com.azure.resourcemanager.postgresqlflexibleserver.models.ServerPublicNetworkAccessState;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Sku;
import com.azure.resourcemanager.postgresqlflexibleserver.models.SkuTier;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Storage;
import com.azure.resourcemanager.postgresqlflexibleserver.models.StorageAutoGrow;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Servers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersClusterCreate.json
     */
    /**
     * Sample code: Create a new elastic cluster.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        createANewElasticCluster(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().define("exampleserver").withRegion("eastus").withExistingResourceGroup("exampleresourcegroup")
            .withSku(new Sku().withName("Standard_D4ds_v5").withTier(SkuTier.GENERAL_PURPOSE))
            .withAdministratorLogin("examplelogin").withAdministratorLoginPassword("examplepassword")
            .withVersion(PostgresMajorVersion.ONE_SIX)
            .withStorage(new Storage().withStorageSizeGB(256).withAutoGrow(StorageAutoGrow.DISABLED)
                .withTier(AzureManagedDiskPerformanceTier.P15))
            .withBackup(
                new Backup().withBackupRetentionDays(7).withGeoRedundantBackup(GeographicallyRedundantBackup.DISABLED))
            .withNetwork(new Network().withPublicNetworkAccess(ServerPublicNetworkAccessState.DISABLED))
            .withHighAvailability(new HighAvailability().withMode(HighAvailabilityMode.DISABLED))
            .withCreateMode(CreateMode.CREATE)
            .withCluster(new Cluster().withClusterSize(2).withDefaultDatabaseName("clusterdb")).create();
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
