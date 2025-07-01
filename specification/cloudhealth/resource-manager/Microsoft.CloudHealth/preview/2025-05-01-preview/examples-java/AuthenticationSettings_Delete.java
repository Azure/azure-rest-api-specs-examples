
/**
 * Samples for AuthenticationSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/AuthenticationSettings_Delete.json
     */
    /**
     * Sample code: AuthenticationSettings_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void authenticationSettingsDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.authenticationSettings().deleteWithResponse("my-resource-group", "my-health-model", "my-auth-setting",
            com.azure.core.util.Context.NONE);
    }
}
