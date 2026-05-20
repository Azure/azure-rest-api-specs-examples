
import com.azure.resourcemanager.storage.models.ActiveDirectoryProperties;
import com.azure.resourcemanager.storage.models.ActiveDirectoryPropertiesAccountType;
import com.azure.resourcemanager.storage.models.AzureFilesIdentityBasedAuthentication;
import com.azure.resourcemanager.storage.models.DirectoryServiceOptions;
import com.azure.resourcemanager.storage.models.StorageAccountUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountEnableAD.json
     */
    /**
     * Sample code: StorageAccountEnableAD.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountEnableAD(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().updateWithResponse("res9407", "sto8596",
            new StorageAccountUpdateParameters().withAzureFilesIdentityBasedAuthentication(
                new AzureFilesIdentityBasedAuthentication().withDirectoryServiceOptions(DirectoryServiceOptions.AD)
                    .withActiveDirectoryProperties(new ActiveDirectoryProperties().withDomainName("adtest.com")
                        .withNetBiosDomainName("adtest.com").withForestName("adtest.com")
                        .withDomainGuid("aebfc118-9fa9-4732-a21f-d98e41a77ae1")
                        .withDomainSid("S-1-5-21-2400535526-2334094090-2402026252")
                        .withAzureStorageSid("S-1-5-21-2400535526-2334094090-2402026252-0012")
                        .withSamAccountName("sam12498").withAccountType(ActiveDirectoryPropertiesAccountType.USER))),
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
