
/**
 * Samples for SensitivitySettings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-02-15-preview/examples/
     * SensitivitySettings/GetSensitivitySettingsList_example.json
     */
    /**
     * Sample code: Get sensitivity settings list.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSensitivitySettingsList(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.sensitivitySettings().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
