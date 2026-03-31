
/**
 * Samples for WebApps GetAppSettingsKeyVaultReferences.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetKeyVaultReferencesForAppSettings.json
     */
    /**
     * Sample code: Get Azure Key Vault references for app settings.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAzureKeyVaultReferencesForAppSettings(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getAppSettingsKeyVaultReferences("testrg123", "testc6282",
            com.azure.core.util.Context.NONE);
    }
}
