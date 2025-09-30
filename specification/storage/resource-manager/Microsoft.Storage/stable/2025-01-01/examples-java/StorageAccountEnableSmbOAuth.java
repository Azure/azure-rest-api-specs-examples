
import com.azure.resourcemanager.storage.models.AzureFilesIdentityBasedAuthentication;
import com.azure.resourcemanager.storage.models.DirectoryServiceOptions;
import com.azure.resourcemanager.storage.models.SmbOAuthSettings;
import com.azure.resourcemanager.storage.models.StorageAccountUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/StorageAccountEnableSmbOAuth.
     * json
     */
    /**
     * Sample code: StorageAccountEnableSmbOAuth.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountEnableSmbOAuth(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().updateWithResponse("res9407", "sto8596",
            new StorageAccountUpdateParameters().withAzureFilesIdentityBasedAuthentication(
                new AzureFilesIdentityBasedAuthentication().withDirectoryServiceOptions(DirectoryServiceOptions.NONE)
                    .withSmbOAuthSettings(new SmbOAuthSettings().withIsSmbOAuthEnabled(true))),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
