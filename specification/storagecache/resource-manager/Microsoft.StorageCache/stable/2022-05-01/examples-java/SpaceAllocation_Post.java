import com.azure.core.util.Context;
import com.azure.resourcemanager.storagecache.models.StorageTargetSpaceAllocation;
import java.util.Arrays;

/** Samples for Caches SpaceAllocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/SpaceAllocation_Post.json
     */
    /**
     * Sample code: SpaceAllocation_Post.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void spaceAllocationPost(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager
            .caches()
            .spaceAllocation(
                "scgroup",
                "sc1",
                Arrays
                    .asList(
                        new StorageTargetSpaceAllocation().withName("st1").withAllocationPercentage(25),
                        new StorageTargetSpaceAllocation().withName("st2").withAllocationPercentage(50),
                        new StorageTargetSpaceAllocation().withName("st3").withAllocationPercentage(25)),
                Context.NONE);
    }
}
