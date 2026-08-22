
import com.azure.resourcemanager.storagesync.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.storagesync.models.Type;

/**
 * Samples for StorageSyncServices CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/StorageSyncServiceCheckNameAvailability_AlreadyExists.json
     */
    /**
     * Sample code: StorageSyncServiceCheckNameAvailability_AlreadyExists.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void storageSyncServiceCheckNameAvailabilityAlreadyExists(
        com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.storageSyncServices()
            .checkNameAvailabilityWithResponse("westus", new CheckNameAvailabilityParameters()
                .withName("newstoragesyncservicename").withType(Type.MICROSOFT_STORAGE_SYNC_STORAGE_SYNC_SERVICES),
                com.azure.core.util.Context.NONE);
    }
}
