import com.azure.resourcemanager.batch.models.ApplicationPackageReference;
import com.azure.resourcemanager.batch.models.CertificateReference;
import com.azure.resourcemanager.batch.models.CertificateStoreLocation;
import com.azure.resourcemanager.batch.models.MetadataItem;
import com.azure.resourcemanager.batch.models.NodeCommunicationMode;
import com.azure.resourcemanager.batch.models.Pool;
import java.util.Arrays;

/** Samples for Pool Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolUpdate_OtherProperties.json
     */
    /**
     * Sample code: UpdatePool - Other Properties.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void updatePoolOtherProperties(com.azure.resourcemanager.batch.BatchManager manager) {
        Pool resource =
            manager
                .pools()
                .getWithResponse(
                    "default-azurebatch-japaneast", "sampleacct", "testpool", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withMetadata(Arrays.asList(new MetadataItem().withName("key1").withValue("value1")))
            .withCertificates(
                Arrays
                    .asList(
                        new CertificateReference()
                            .withId(
                                "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/certificates/sha1-1234567")
                            .withStoreLocation(CertificateStoreLocation.LOCAL_MACHINE)
                            .withStoreName("MY")))
            .withApplicationPackages(
                Arrays
                    .asList(
                        new ApplicationPackageReference()
                            .withId(
                                "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
                        new ApplicationPackageReference()
                            .withId(
                                "/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_5678")
                            .withVersion("1.0")))
            .withTargetNodeCommunicationMode(NodeCommunicationMode.SIMPLIFIED)
            .apply();
    }
}
