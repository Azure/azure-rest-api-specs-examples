
import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;
import com.azure.resourcemanager.storage.models.ImmutableStorageWithVersioning;

/**
 * Samples for BlobContainers Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * BlobContainersPutObjectLevelWorm.json
     */
    /**
     * Sample code: PutContainerWithObjectLevelWorm.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putContainerWithObjectLevelWorm(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobContainers()
            .createWithResponse("res3376", "sto328", "container6185", new BlobContainerInner()
                .withImmutableStorageWithVersioning(new ImmutableStorageWithVersioning().withEnabled(true)),
                com.azure.core.util.Context.NONE);
    }
}
