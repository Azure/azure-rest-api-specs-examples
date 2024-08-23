
/**
 * Samples for ManagedDatabaseSensitivityLabels ListRecommendedByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseSensitivityLabelsListByDatabaseRecommended.json
     */
    /**
     * Sample code: Gets the recommended sensitivity labels of a given database in a managed database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheRecommendedSensitivityLabelsOfAGivenDatabaseInAManagedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSensitivityLabels().listRecommendedByDatabase(
            "myRG", "myManagedInstanceName", "myDatabase", null, null, null, com.azure.core.util.Context.NONE);
    }
}
