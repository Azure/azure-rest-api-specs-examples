
import com.azure.resourcemanager.storage.models.ContextCacheCheckNameAvailabilityParameters;

/**
 * Samples for ContextCaches CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/StorageContextCacheCRUD/ContextCaches_CheckNameAvailability.json
     */
    /**
     * Sample code: ContextCacheCheckNameAvailability.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void contextCacheCheckNameAvailability(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCaches().checkNameAvailabilityWithResponse(
            new ContextCacheCheckNameAvailabilityParameters().withName("testcontextcache"),
            com.azure.core.util.Context.NONE);
    }
}
