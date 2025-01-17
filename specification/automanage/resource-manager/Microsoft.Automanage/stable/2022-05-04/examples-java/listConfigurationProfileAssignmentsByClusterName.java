
/**
 * Samples for ConfigurationProfileAssignments ListByClusterName.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/
     * listConfigurationProfileAssignmentsByClusterName.json
     */
    /**
     * Sample code: List configuration profile assignments by resourceGroup and cluster.
     * 
     * @param manager Entry point to AutomanageManager.
     */
    public static void listConfigurationProfileAssignmentsByResourceGroupAndCluster(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.configurationProfileAssignments().listByClusterName("myResourceGroupName", "myClusterName",
            com.azure.core.util.Context.NONE);
    }
}
