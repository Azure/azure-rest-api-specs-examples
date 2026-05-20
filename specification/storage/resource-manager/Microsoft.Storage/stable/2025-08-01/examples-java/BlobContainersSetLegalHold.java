
import com.azure.resourcemanager.storage.fluent.models.LegalHoldInner;
import java.util.Arrays;

/**
 * Samples for BlobContainers SetLegalHold.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/BlobContainersSetLegalHold.json
     */
    /**
     * Sample code: SetLegalHoldContainers.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void setLegalHoldContainers(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().setLegalHoldWithResponse("res4303", "sto7280", "container8723",
            new LegalHoldInner().withTags(Arrays.asList("tag1", "tag2", "tag3")), com.azure.core.util.Context.NONE);
    }
}
