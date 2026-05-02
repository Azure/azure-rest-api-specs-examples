
import com.azure.resourcemanager.storagecache.models.PrimingJobIdParameter;

/**
 * Samples for Caches StopPrimingJob.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/StopPrimingJob.json
     */
    /**
     * Sample code: StopPrimingJob.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void stopPrimingJob(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().stopPrimingJob("scgroup", "sc1",
            new PrimingJobIdParameter().withPrimingJobId("00000000000_0000000000"), com.azure.core.util.Context.NONE);
    }
}
