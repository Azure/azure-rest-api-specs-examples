
import com.azure.resourcemanager.netapp.models.Bucket;
import com.azure.resourcemanager.netapp.models.BucketPatchPermissions;
import com.azure.resourcemanager.netapp.models.BucketServerPatchProperties;
import com.azure.resourcemanager.netapp.models.OnCertificateConflictAction;

/**
 * Samples for Buckets Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/Buckets_Update.json
     */
    /**
     * Sample code: Buckets_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void bucketsUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        Bucket resource = manager.buckets()
            .getWithResponse("myRG", "account1", "pool1", "volume1", "bucket1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withServer(new BucketServerPatchProperties().withFqdn("fullyqualified.domainname.com")
            .withCertificateObject("<REDACTED>").withOnCertificateConflictAction(OnCertificateConflictAction.UPDATE))
            .withPermissions(BucketPatchPermissions.READ_WRITE).apply();
    }
}
