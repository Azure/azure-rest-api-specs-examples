
import com.azure.resourcemanager.storage.models.StorageAccountExpand;

/**
 * Samples for StorageAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/
     * StorageAccountGetPropertiesGeoReplicationStatscanFailoverTrue.json
     */
    /**
     * Sample code: StorageAccountGetPropertiesGeoReplicationStatscanFailoverTrue.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountGetPropertiesGeoReplicationStatscanFailoverTrue(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().getByResourceGroupWithResponse("res9407",
            "sto8596", StorageAccountExpand.GEO_REPLICATION_STATS, com.azure.core.util.Context.NONE);
    }
}
