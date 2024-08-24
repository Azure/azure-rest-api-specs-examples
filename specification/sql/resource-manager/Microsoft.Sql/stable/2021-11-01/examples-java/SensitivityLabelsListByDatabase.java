
/**
 * Samples for SensitivityLabels ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SensitivityLabelsListByDatabase.json
     */
    /**
     * Sample code: Gets the current and recommended sensitivity labels of a given database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheCurrentAndRecommendedSensitivityLabelsOfAGivenDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSensitivityLabels().listByDatabase("myRG", "myServer",
            "myDatabase", null, com.azure.core.util.Context.NONE);
    }
}
