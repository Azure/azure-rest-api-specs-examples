
/**
 * Samples for WebApps GetAppSettingKeyVaultReferenceSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetKeyVaultReferencesForAppSettingSlot.json
     */
    /**
     * Sample code: Get Azure Key Vault slot app setting reference.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAzureKeyVaultSlotAppSettingReference(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getAppSettingKeyVaultReferenceSlotWithResponse("testrg123", "testc6282",
            "setting", "stage", com.azure.core.util.Context.NONE);
    }
}
