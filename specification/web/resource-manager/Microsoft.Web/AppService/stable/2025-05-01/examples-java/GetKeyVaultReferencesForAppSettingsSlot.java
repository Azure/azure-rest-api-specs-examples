
/**
 * Samples for WebApps GetAppSettingsKeyVaultReferencesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetKeyVaultReferencesForAppSettingsSlot.json
     */
    /**
     * Sample code: Get Azure Key Vault references for app settings for slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getAzureKeyVaultReferencesForAppSettingsForSlot(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getAppSettingsKeyVaultReferencesSlot("testrg123", "testc6282", "stage",
            com.azure.core.util.Context.NONE);
    }
}
