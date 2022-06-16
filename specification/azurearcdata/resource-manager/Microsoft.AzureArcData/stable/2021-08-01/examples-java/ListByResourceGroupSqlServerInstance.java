import com.azure.core.util.Context;

/** Samples for SqlServerInstances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/ListByResourceGroupSqlServerInstance.json
     */
    /**
     * Sample code: Gets all SQL Server Instance in a resource group.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void getsAllSQLServerInstanceInAResourceGroup(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.sqlServerInstances().listByResourceGroup("testrg", Context.NONE);
    }
}
