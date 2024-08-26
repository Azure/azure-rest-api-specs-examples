
/**
 * Samples for SecuritySettings ListByClusters.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListSecuritySettingsByCluster.json
     */
    /**
     * Sample code: List Security Settings.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listSecuritySettings(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.securitySettings().listByClusters("test-rg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
