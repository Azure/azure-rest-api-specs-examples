
import com.azure.resourcemanager.storage.models.StorageAccountExpand;

/**
 * Samples for StorageAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountGetPropertiesGeoReplicationStatscanFailoverTrue.json
     */
    /**
     * Sample code: StorageAccountGetPropertiesGeoReplicationStatscanFailoverTrue.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountGetPropertiesGeoReplicationStatscanFailoverTrue(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().getByResourceGroupWithResponse("res9407", "sto8596",
            StorageAccountExpand.GEO_REPLICATION_STATS, com.azure.core.util.Context.NONE);
    }
}
