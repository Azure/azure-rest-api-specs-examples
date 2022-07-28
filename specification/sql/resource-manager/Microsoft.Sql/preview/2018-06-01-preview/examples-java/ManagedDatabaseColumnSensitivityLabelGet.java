import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.SensitivityLabelSource;

/** Samples for ManagedDatabaseSensitivityLabels Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ManagedDatabaseColumnSensitivityLabelGet.json
     */
    /**
     * Sample code: Gets the sensitivity label of a given column in a managed database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheSensitivityLabelOfAGivenColumnInAManagedDatabase(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedDatabaseSensitivityLabels()
            .getWithResponse(
                "myRG",
                "myManagedInstanceName",
                "myDatabase",
                "dbo",
                "myTable",
                "myColumn",
                SensitivityLabelSource.CURRENT,
                Context.NONE);
    }
}
