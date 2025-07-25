
import com.azure.resourcemanager.appservice.fluent.models.AzureStoragePropertyDictionaryResourceInner;
import com.azure.resourcemanager.appservice.models.AzureStorageInfoValue;
import com.azure.resourcemanager.appservice.models.AzureStorageType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for WebApps UpdateAzureStorageAccounts.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdateAzureStorageAccounts.json
     */
    /**
     * Sample code: Update Azure Storage Accounts.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAzureStorageAccounts(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().updateAzureStorageAccountsWithResponse("testrg123",
            "sitef6141",
            new AzureStoragePropertyDictionaryResourceInner().withProperties(mapOf("account1",
                new AzureStorageInfoValue().withType(AzureStorageType.AZURE_FILES).withAccountName("testsa")
                    .withShareName("web").withAccessKey("fakeTokenPlaceholder").withMountPath("/mounts/a/files"))),
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
