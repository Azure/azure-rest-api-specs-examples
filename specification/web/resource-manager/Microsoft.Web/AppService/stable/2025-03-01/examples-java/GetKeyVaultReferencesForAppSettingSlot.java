
/**
 * Samples for WebApps GetAppSettingKeyVaultReferenceSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetKeyVaultReferencesForAppSettingSlot.json
     */
    /**
     * Sample code: Get Azure Key Vault slot app setting reference.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureKeyVaultSlotAppSettingReference(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getAppSettingKeyVaultReferenceSlotWithResponse(
            "testrg123", "testc6282", "setting", "stage", com.azure.core.util.Context.NONE);
    }
}
