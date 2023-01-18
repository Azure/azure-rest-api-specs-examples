/** Samples for ManagementConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementConfigurationDelete.json
     */
    /**
     * Sample code: ManagementConfigurationDelete.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void managementConfigurationDelete(
        com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager
            .managementConfigurations()
            .deleteByResourceGroupWithResponse("rg1", "managementConfigurationName", com.azure.core.util.Context.NONE);
    }
}
