
import com.azure.resourcemanager.batch.models.ApplicationPackageReference;
import com.azure.resourcemanager.batch.models.MetadataItem;
import com.azure.resourcemanager.batch.models.Pool;
import java.util.Arrays;

/**
 * Samples for Pool Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PoolUpdate_OtherProperties.json
     */
    /**
     * Sample code: UpdatePool - Other Properties.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void updatePoolOtherProperties(com.azure.resourcemanager.batch.BatchManager manager) {
        Pool resource = manager.pools()
            .getWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withMetadata(Arrays.asList(new MetadataItem().withName("key1").withValue("value1")))
            .withApplicationPackages(Arrays.asList(new ApplicationPackageReference().withId(
                "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_1234"),
                new ApplicationPackageReference().withId(
                    "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Batch/batchAccounts/sampleacct/pools/testpool/applications/app_5678")
                    .withVersion("1.0")))
            .apply();
    }
}
