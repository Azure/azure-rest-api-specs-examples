
/**
 * Samples for WebApps GetAuthSettingsV2WithoutSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetAuthSettingsV2WithoutSecrets.json
     */
    /**
     * Sample code: List Auth Settings without Secrets.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAuthSettingsWithoutSecrets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getAuthSettingsV2WithoutSecretsWithResponse("testrg123",
            "sitef6141", com.azure.core.util.Context.NONE);
    }
}
