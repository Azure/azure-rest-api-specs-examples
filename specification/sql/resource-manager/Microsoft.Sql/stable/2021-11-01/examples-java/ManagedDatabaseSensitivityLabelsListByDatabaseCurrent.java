
import com.azure.core.util.Context;

/** Samples for ManagedDatabaseSensitivityLabels ListCurrentByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseSensitivityLabelsListByDatabaseCurrent.json
     */
    /**
     * Sample code: Gets the current sensitivity labels of a given database in a managed database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheCurrentSensitivityLabelsOfAGivenDatabaseInAManagedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSensitivityLabels().listCurrentByDatabase("myRG",
            "myManagedInstanceName", "myDatabase", null, null, null, Context.NONE);
    }
}
