import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.FileShareInner;

/** Samples for FileShares Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/FileSharesPut.json
     */
    /**
     * Sample code: PutShares.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putShares(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .createWithResponse("res3376", "sto328", "share6185", new FileShareInner(), null, Context.NONE);
    }
}
