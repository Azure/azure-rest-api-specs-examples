
import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;
import com.azure.resourcemanager.storage.models.BlobAccessPointConfigurationConnection;

/**
 * Samples for BlobContainers Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/BlobContainersPutBlobAccessPointConfiguration.json
     */
    /**
     * Sample code: PutContainerWithBlobAccessPointConfiguration.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        putContainerWithBlobAccessPointConfiguration(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().createWithResponse("res3376", "sto328", "container6185",
            new BlobContainerInner().withBlobAccessPointConfiguration(
                new BlobAccessPointConfigurationConnection().withBlobAccessPointConfigurationName("myAccessPointConfig")
                    .withBlobAccessPointConfigurationUniqueId("00000000-0000-0000-0000-000000000000")),
            com.azure.core.util.Context.NONE);
    }
}
