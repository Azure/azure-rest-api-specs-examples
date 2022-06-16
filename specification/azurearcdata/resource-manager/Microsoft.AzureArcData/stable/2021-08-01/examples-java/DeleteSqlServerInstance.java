import com.azure.core.util.Context;

/** Samples for SqlServerInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/DeleteSqlServerInstance.json
     */
    /**
     * Sample code: Delete a SQL Server Instance.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void deleteASQLServerInstance(com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.sqlServerInstances().delete("testrg", "testsqlServerInstance", Context.NONE);
    }
}
