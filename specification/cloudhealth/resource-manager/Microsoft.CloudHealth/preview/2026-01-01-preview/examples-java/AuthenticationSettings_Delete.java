
/**
 * Samples for AuthenticationSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/AuthenticationSettings_Delete.json
     */
    /**
     * Sample code: AuthenticationSettings_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void authenticationSettingsDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.authenticationSettings().delete("my-resource-group", "my-health-model", "my-auth-setting",
            com.azure.core.util.Context.NONE);
    }
}
