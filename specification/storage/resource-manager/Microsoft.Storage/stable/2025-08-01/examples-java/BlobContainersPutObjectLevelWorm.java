
import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;
import com.azure.resourcemanager.storage.models.ImmutableStorageWithVersioning;

/**
 * Samples for BlobContainers Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/BlobContainersPutObjectLevelWorm.json
     */
    /**
     * Sample code: PutContainerWithObjectLevelWorm.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putContainerWithObjectLevelWorm(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers()
            .createWithResponse("res3376", "sto328", "container6185", new BlobContainerInner()
                .withImmutableStorageWithVersioning(new ImmutableStorageWithVersioning().withEnabled(true)),
                com.azure.core.util.Context.NONE);
    }
}
