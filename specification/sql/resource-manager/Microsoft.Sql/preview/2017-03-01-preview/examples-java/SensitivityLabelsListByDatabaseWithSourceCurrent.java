import com.azure.core.util.Context;

/** Samples for SensitivityLabels ListCurrentByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/SensitivityLabelsListByDatabaseWithSourceCurrent.json
     */
    /**
     * Sample code: Gets the current sensitivity labels of a given database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheCurrentSensitivityLabelsOfAGivenDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getSensitivityLabels()
            .listCurrentByDatabase("myRG", "myServer", "myDatabase", null, Context.NONE);
    }
}
