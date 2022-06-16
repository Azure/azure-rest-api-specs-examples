import com.azure.core.util.Context;

/** Samples for WebApps GetAppSettingsKeyVaultReferencesSlot. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSettingsSlot.json
     */
    /**
     * Sample code: Get Azure Key Vault references for app settings for slot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureKeyVaultReferencesForAppSettingsForSlot(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getAppSettingsKeyVaultReferencesSlot("testrg123", "testc6282", "stage", Context.NONE);
    }
}
