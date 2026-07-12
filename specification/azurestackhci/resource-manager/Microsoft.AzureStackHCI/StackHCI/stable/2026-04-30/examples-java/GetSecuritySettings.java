
/**
 * Samples for SecuritySettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/GetSecuritySettings.json
     */
    /**
     * Sample code: Get Security Settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getSecuritySettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.securitySettings().getWithResponse("test-rg", "myCluster", "default", com.azure.core.util.Context.NONE);
    }
}
