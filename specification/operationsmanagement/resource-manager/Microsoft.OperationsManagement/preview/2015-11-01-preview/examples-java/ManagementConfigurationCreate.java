/** Samples for ManagementConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementConfigurationCreate.json
     */
    /**
     * Sample code: ManagementConfigurationCreate.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void managementConfigurationCreate(
        com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager
            .managementConfigurations()
            .define("managementConfiguration1")
            .withExistingResourceGroup("rg1")
            .withRegion("East US")
            .create();
    }
}
