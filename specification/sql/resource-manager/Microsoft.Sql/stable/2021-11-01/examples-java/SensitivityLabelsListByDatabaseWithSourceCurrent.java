
import com.azure.core.util.Context;

/** Samples for SensitivityLabels ListCurrentByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * SensitivityLabelsListByDatabaseWithSourceCurrent.json
     */
    /**
     * Sample code: Gets the current sensitivity labels of a given database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsTheCurrentSensitivityLabelsOfAGivenDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSensitivityLabels().listCurrentByDatabase("myRG", "myServer",
            "myDatabase", null, null, null, Context.NONE);
    }
}
