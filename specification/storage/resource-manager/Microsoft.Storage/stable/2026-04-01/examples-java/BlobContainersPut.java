
import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;

/**
 * Samples for BlobContainers Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobContainersPut.json
     */
    /**
     * Sample code: PutContainers.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putContainers(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().createWithResponse("res3376", "sto328", "container6185",
            new BlobContainerInner(), com.azure.core.util.Context.NONE);
    }
}
