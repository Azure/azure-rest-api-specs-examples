import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;

/** Samples for BlobContainers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobContainersPut.json
     */
    /**
     * Sample code: PutContainers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putContainers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .createWithResponse(
                "res3376", "sto328", "container6185", new BlobContainerInner(), com.azure.core.util.Context.NONE);
    }
}
