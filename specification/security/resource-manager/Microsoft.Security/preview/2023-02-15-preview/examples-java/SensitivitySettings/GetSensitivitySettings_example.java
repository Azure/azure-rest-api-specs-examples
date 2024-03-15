
/**
 * Samples for ResourceProvider GetSensitivitySettings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-02-15-preview/examples/
     * SensitivitySettings/GetSensitivitySettings_example.json
     */
    /**
     * Sample code: Get sensitivity settings.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSensitivitySettings(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.resourceProviders().getSensitivitySettingsWithResponse(com.azure.core.util.Context.NONE);
    }
}
