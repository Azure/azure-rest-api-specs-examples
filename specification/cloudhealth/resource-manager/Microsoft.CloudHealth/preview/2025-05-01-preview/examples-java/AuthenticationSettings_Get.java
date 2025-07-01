
/**
 * Samples for AuthenticationSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/AuthenticationSettings_Get.json
     */
    /**
     * Sample code: AuthenticationSettings_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void authenticationSettingsGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.authenticationSettings().getWithResponse("my-resource-group", "my-health-model", "my-auth-setting",
            com.azure.core.util.Context.NONE);
    }
}
