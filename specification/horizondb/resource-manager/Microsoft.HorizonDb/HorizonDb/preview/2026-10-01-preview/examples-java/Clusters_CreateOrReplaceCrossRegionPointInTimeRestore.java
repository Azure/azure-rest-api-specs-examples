
import com.azure.resourcemanager.horizondb.models.BackupProperties;
import com.azure.resourcemanager.horizondb.models.ClusterCreateMode;
import com.azure.resourcemanager.horizondb.models.ClusterProperties;
import com.azure.resourcemanager.horizondb.models.GeoRedundancyState;
import com.azure.resourcemanager.horizondb.models.Network;
import com.azure.resourcemanager.horizondb.models.PublicNetworkAccessState;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_CreateOrReplaceCrossRegionPointInTimeRestore.json
     */
    /**
     * Sample code: Restore a HorizonDB cluster across Azure regions from the latest available backup.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void restoreAHorizonDBClusterAcrossAzureRegionsFromTheLatestAvailableBackup(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().define("cross-region-restored-cluster").withRegion("westus2")
            .withExistingResourceGroup("example-resource-group")
            .withProperties(new ClusterProperties().withCreationMode(ClusterCreateMode.POINT_IN_TIME_RESTORE)
                .withPointInTimeUtc(OffsetDateTime.parse("2026-07-27T12:00:00Z"))
                .withCreationSourceClusterResourceId(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/source-cluster")
                .withCreationSourceClusterLocation("eastus2")
                .withNetwork(new Network().withPublicAccess(PublicNetworkAccessState.DISABLED))
                .withBackup(new BackupProperties().withGeoRedundancy(GeoRedundancyState.ENABLED).withRetentionDays(14)))
            .create();
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
