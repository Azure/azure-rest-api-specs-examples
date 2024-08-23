
/**
 * Samples for ManagedDatabaseSensitivityLabels Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseColumnSensitivityLabelDelete.json
     */
    /**
     * Sample code: Deletes the sensitivity label of a given column in a managed database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesTheSensitivityLabelOfAGivenColumnInAManagedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSensitivityLabels().deleteWithResponse("myRG",
            "myManagedInstanceName", "myDatabase", "dbo", "myTable", "myColumn", com.azure.core.util.Context.NONE);
    }
}
