
import com.azure.core.util.Context;

/** Samples for SensitivityLabels ListRecommendedByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * SensitivityLabelsListByDatabaseWithSourceRecommended.json
     */
    /**
     * Sample code: Gets the recommended sensitivity labels of a given database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsTheRecommendedSensitivityLabelsOfAGivenDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSensitivityLabels().listRecommendedByDatabase("myRG",
            "myServer", "myDatabase", null, null, null, Context.NONE);
    }
}
