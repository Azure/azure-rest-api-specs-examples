
/**
 * Samples for ArcSettings ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListArcSettingsByCluster.json
     */
    /**
     * Sample code: List ArcSetting resources by HCI Cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listArcSettingResourcesByHCICluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().listByCluster("test-rg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
