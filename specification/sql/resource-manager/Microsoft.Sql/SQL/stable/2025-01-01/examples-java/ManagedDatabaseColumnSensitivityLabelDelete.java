
import com.azure.resourcemanager.sql.models.CurrentSensitivityLabelSource;

/**
 * Samples for ManagedDatabaseSensitivityLabels Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseColumnSensitivityLabelDelete.json
     */
    /**
     * Sample code: Deletes the sensitivity label of a given column in a managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deletesTheSensitivityLabelOfAGivenColumnInAManagedDatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSensitivityLabels().deleteWithResponse("myRG",
            "myManagedInstanceName", "myDatabase", "dbo", "myTable", "myColumn", CurrentSensitivityLabelSource.CURRENT,
            com.azure.core.util.Context.NONE);
    }
}
