
import com.azure.resourcemanager.netapp.models.CheckElasticVolumeFilePathAvailabilityRequest;

/**
 * Samples for ElasticCapacityPools CheckVolumeFilePathAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-15-preview/ElasticCapacityPools_CheckVolumeFilePathAvailability.json
     */
    /**
     * Sample code: ElasticCapacityPools_CheckVolumeFilePathAvailability.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticCapacityPoolsCheckVolumeFilePathAvailability(
        com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticCapacityPools().checkVolumeFilePathAvailabilityWithResponse("myRG", "account1", "pool1",
            new CheckElasticVolumeFilePathAvailabilityRequest().withFilePath("my-exact-filepath"),
            com.azure.core.util.Context.NONE);
    }
}
