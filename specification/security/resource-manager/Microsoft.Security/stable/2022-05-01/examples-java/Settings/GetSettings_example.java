
/**
 * Samples for Settings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2022-05-01/examples/Settings/
     * GetSettings_example.json
     */
    /**
     * Sample code: Get settings of subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSettingsOfSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.settings().list(com.azure.core.util.Context.NONE);
    }
}
