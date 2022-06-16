import com.azure.core.util.Context;

/** Samples for SqlServerInstances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/GetSqlServerInstance.json
     */
    /**
     * Sample code: Updates a SQL Server Instance tags.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void updatesASQLServerInstanceTags(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.sqlServerInstances().getByResourceGroupWithResponse("testrg", "testsqlServerInstance", Context.NONE);
    }
}
