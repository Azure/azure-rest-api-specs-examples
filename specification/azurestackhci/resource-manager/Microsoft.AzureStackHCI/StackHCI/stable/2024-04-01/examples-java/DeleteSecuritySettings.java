
/**
 * Samples for SecuritySettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * DeleteSecuritySettings.json
     */
    /**
     * Sample code: Delete Security Settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteSecuritySettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.securitySettings().delete("test-rg", "myCluster", "default", com.azure.core.util.Context.NONE);
    }
}
