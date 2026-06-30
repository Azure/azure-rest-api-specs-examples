
import com.azure.resourcemanager.netapp.models.BucketCredentialsExpiry;

/**
 * Samples for Buckets GenerateAkvCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/Buckets_GenerateAkvCredentials.json
     */
    /**
     * Sample code: Buckets_GenerateAkvCredentials.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void bucketsGenerateAkvCredentials(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.buckets().generateAkvCredentials("myRG", "account1", "pool1", "volume1", "bucket1",
            new BucketCredentialsExpiry().withKeyPairExpiryDays(3), com.azure.core.util.Context.NONE);
    }
}
