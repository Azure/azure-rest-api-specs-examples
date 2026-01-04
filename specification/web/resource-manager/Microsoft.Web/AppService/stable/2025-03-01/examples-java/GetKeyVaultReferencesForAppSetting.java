
/**
 * Samples for WebApps GetAppSettingKeyVaultReference.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetKeyVaultReferencesForAppSetting.json
     */
    /**
     * Sample code: Get Azure Key Vault app setting reference.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAzureKeyVaultAppSettingReference(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getAppSettingKeyVaultReferenceWithResponse("testrg123",
            "testc6282", "setting", com.azure.core.util.Context.NONE);
    }
}
