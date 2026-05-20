
import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;

/**
 * Samples for BlobContainers Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/BlobContainersPutDefaultEncryptionScope.json
     */
    /**
     * Sample code: PutContainerWithDefaultEncryptionScope.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        putContainerWithDefaultEncryptionScope(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().createWithResponse(
            "res3376", "sto328", "container6185", new BlobContainerInner()
                .withDefaultEncryptionScope("encryptionscope185").withDenyEncryptionScopeOverride(true),
            com.azure.core.util.Context.NONE);
    }
}
