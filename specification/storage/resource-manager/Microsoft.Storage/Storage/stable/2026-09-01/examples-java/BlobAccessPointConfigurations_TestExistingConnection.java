
import com.azure.resourcemanager.storage.models.BlobAccessPointConnectionTestRequest;

/**
 * Samples for BlobAccessPointConfigurations TestExistingConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/BlobAccessPointConfigurations_TestExistingConnection.json
     */
    /**
     * Sample code: BlobAccessPointConfigurationTestExistingConnection.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        blobAccessPointConfigurationTestExistingConnection(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobAccessPointConfigurations().testExistingConnection("testrg",
            "teststorageaccount", "testaccesspointconfig",
            new BlobAccessPointConnectionTestRequest().withUniqueId("11111111-2222-3333-4444-555555555555"),
            com.azure.core.util.Context.NONE);
    }
}
