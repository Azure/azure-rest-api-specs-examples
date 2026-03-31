
/**
 * Samples for WebApps GetAppSettingKeyVaultReference.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetKeyVaultReferencesForAppSetting.json
     */
    /**
     * Sample code: Get Azure Key Vault app setting reference.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAzureKeyVaultAppSettingReference(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getAppSettingKeyVaultReferenceWithResponse("testrg123", "testc6282",
            "setting", com.azure.core.util.Context.NONE);
    }
}
