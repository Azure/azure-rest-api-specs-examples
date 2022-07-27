import com.azure.core.util.Context;

/** Samples for WebApps GetAppSettingsKeyVaultReferences. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetKeyVaultReferencesForAppSettings.json
     */
    /**
     * Sample code: Get Azure Key Vault references for app settings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureKeyVaultReferencesForAppSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getAppSettingsKeyVaultReferences("testrg123", "testc6282", Context.NONE);
    }
}
