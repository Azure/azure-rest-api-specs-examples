import com.azure.core.util.Context;
import com.azure.resourcemanager.storagecache.models.PrimingJobIdParameter;

/** Samples for Caches ResumePrimingJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/ResumePrimingJob.json
     */
    /**
     * Sample code: ResumePrimingJob.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void resumePrimingJob(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager
            .caches()
            .resumePrimingJob(
                "scgroup", "sc1", new PrimingJobIdParameter().withPrimingJobId("00000000000_0000000000"), Context.NONE);
    }
}
