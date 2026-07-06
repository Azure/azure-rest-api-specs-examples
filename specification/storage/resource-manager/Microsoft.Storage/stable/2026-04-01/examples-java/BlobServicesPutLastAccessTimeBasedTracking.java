
import com.azure.resourcemanager.storage.fluent.models.BlobServicePropertiesInner;
import com.azure.resourcemanager.storage.models.LastAccessTimeTrackingPolicy;
import com.azure.resourcemanager.storage.models.Name;
import java.util.Arrays;

/**
 * Samples for BlobServices SetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobServicesPutLastAccessTimeBasedTracking.json
     */
    /**
     * Sample code: BlobServicesPutLastAccessTimeBasedTracking.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        blobServicesPutLastAccessTimeBasedTracking(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobServices().setServicePropertiesWithResponse("res4410", "sto8607",
            new BlobServicePropertiesInner().withLastAccessTimeTrackingPolicy(
                new LastAccessTimeTrackingPolicy().withEnable(true).withName(Name.ACCESS_TIME_TRACKING)
                    .withTrackingGranularityInDays(1).withBlobType(Arrays.asList("blockBlob"))),
            com.azure.core.util.Context.NONE);
    }
}
