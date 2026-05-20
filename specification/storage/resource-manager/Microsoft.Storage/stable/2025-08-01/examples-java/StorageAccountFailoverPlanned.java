
import com.azure.resourcemanager.storage.models.FailoverType;

/**
 * Samples for StorageAccounts Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountFailoverPlanned.json
     */
    /**
     * Sample code: StorageAccountFailoverPlanned.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountFailoverPlanned(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().failover("res4228", "sto2434", FailoverType.PLANNED,
            com.azure.core.util.Context.NONE);
    }
}
