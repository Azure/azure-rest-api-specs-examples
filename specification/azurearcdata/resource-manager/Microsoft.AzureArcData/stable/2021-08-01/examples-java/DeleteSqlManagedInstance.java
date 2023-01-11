/** Samples for SqlManagedInstances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/DeleteSqlManagedInstance.json
     */
    /**
     * Sample code: Delete a SQL Instance.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void deleteASQLInstance(com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.sqlManagedInstances().delete("testrg", "testsqlManagedInstance", com.azure.core.util.Context.NONE);
    }
}
