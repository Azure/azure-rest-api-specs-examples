
import com.azure.resourcemanager.storage.models.AccountSasParameters;
import com.azure.resourcemanager.storage.models.HttpProtocol;
import com.azure.resourcemanager.storage.models.Permissions;
import com.azure.resourcemanager.storage.models.Services;
import com.azure.resourcemanager.storage.models.SignedResourceTypes;
import java.time.OffsetDateTime;

/**
 * Samples for StorageAccounts ListAccountSas.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/StorageAccountListAccountSAS.
     * json
     */
    /**
     * Sample code: StorageAccountListAccountSAS.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountListAccountSAS(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().listAccountSasWithResponse("res7985",
            "sto8588",
            new AccountSasParameters().withServices(Services.B).withResourceTypes(SignedResourceTypes.S)
                .withPermissions(Permissions.R).withProtocols(HttpProtocol.HTTPS_HTTP)
                .withSharedAccessStartTime(OffsetDateTime.parse("2017-05-24T10:42:03.1567373Z"))
                .withSharedAccessExpiryTime(OffsetDateTime.parse("2017-05-24T11:42:03.1567373Z"))
                .withKeyToSign("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
