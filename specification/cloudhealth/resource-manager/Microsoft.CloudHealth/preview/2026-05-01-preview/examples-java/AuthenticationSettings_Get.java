
/**
 * Samples for AuthenticationSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/AuthenticationSettings_Get.json
     */
    /**
     * Sample code: AuthenticationSettings_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void authenticationSettingsGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.authenticationSettings().getWithResponse("online-store-rg", "online-store", "default-auth",
            com.azure.core.util.Context.NONE);
    }
}
