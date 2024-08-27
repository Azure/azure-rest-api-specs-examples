
/**
 * Samples for SecuritySettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * GetSecuritySettings.json
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
