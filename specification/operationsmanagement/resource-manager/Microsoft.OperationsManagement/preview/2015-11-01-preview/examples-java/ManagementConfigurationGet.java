/** Samples for ManagementConfigurations GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementConfigurationGet.json
     */
    /**
     * Sample code: SolutionGet.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void solutionGet(com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager
            .managementConfigurations()
            .getByResourceGroupWithResponse("rg1", "managementConfigurationName", com.azure.core.util.Context.NONE);
    }
}
