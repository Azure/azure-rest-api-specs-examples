
import com.azure.resourcemanager.storage.models.StorageAccountExpand;

/**
 * Samples for StorageAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountGetPropertiesGeoReplicationStatscanFailoverFalse.json
     */
    /**
     * Sample code: StorageAccountGetPropertiesGeoReplicationStatscanFailoverFalse.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountGetPropertiesGeoReplicationStatscanFailoverFalse(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().getByResourceGroupWithResponse("res9407", "sto8596",
            StorageAccountExpand.GEO_REPLICATION_STATS, com.azure.core.util.Context.NONE);
    }
}
