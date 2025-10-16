
import com.azure.resourcemanager.netapp.models.BucketCredentialsExpiry;

/**
 * Samples for Buckets GenerateCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/
     * Buckets_GenerateCredentials.json
     */
    /**
     * Sample code: Buckets_GenerateCredentials.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void bucketsGenerateCredentials(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.buckets().generateCredentialsWithResponse("myRG", "account1", "pool1", "volume1", "bucket1",
            new BucketCredentialsExpiry().withKeyPairExpiryDays(3), com.azure.core.util.Context.NONE);
    }
}
