import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.BlobContainerInner;

/** Samples for BlobContainers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/BlobContainersPutDefaultEncryptionScope.json
     */
    /**
     * Sample code: PutContainerWithDefaultEncryptionScope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putContainerWithDefaultEncryptionScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .createWithResponse(
                "res3376",
                "sto328",
                "container6185",
                new BlobContainerInner()
                    .withDefaultEncryptionScope("encryptionscope185")
                    .withDenyEncryptionScopeOverride(true),
                Context.NONE);
    }
}
