
/**
 * Samples for WebApps GetAuthSettingsV2WithoutSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetAuthSettingsV2WithoutSecrets.json
     */
    /**
     * Sample code: List Auth Settings without Secrets.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listAuthSettingsWithoutSecrets(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().getAuthSettingsV2WithoutSecretsWithResponse("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
