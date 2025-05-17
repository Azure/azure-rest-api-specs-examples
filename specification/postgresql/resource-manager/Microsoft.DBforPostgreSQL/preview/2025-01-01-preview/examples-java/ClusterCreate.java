
import com.azure.resourcemanager.postgresqlflexibleserver.models.AzureManagedDiskPerformanceTiers;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Backup;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Cluster;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.GeoRedundantBackupEnum;
import com.azure.resourcemanager.postgresqlflexibleserver.models.HighAvailability;
import com.azure.resourcemanager.postgresqlflexibleserver.models.HighAvailabilityMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Network;
import com.azure.resourcemanager.postgresqlflexibleserver.models.ServerPublicNetworkAccessState;
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
     * ClusterCreate.json
     */
    /**
     * Sample code: ClusterCreate.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void clusterCreate(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().define("pgtestcluster").withRegion("westus").withExistingResourceGroup("testrg")
            .withSku(new Sku().withName("Standard_D4ds_v5").withTier(SkuTier.GENERAL_PURPOSE))
            .withAdministratorLogin("login").withAdministratorLoginPassword("Password1")
            .withVersion(ServerVersion.ONE_SIX)
            .withStorage(new Storage().withStorageSizeGB(256).withAutoGrow(StorageAutoGrow.DISABLED)
                .withTier(AzureManagedDiskPerformanceTiers.P15))
            .withBackup(new Backup().withBackupRetentionDays(7).withGeoRedundantBackup(GeoRedundantBackupEnum.DISABLED))
            .withNetwork(new Network().withPublicNetworkAccess(ServerPublicNetworkAccessState.DISABLED))
            .withHighAvailability(new HighAvailability().withMode(HighAvailabilityMode.DISABLED))
            .withCreateMode(CreateMode.CREATE).withCluster(new Cluster().withClusterSize(2)).create();
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
