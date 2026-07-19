
import com.azure.resourcemanager.sql.models.SensitivityLabelSource;

/**
 * Samples for ManagedDatabaseSensitivityLabels Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseColumnSensitivityLabelGet.json
     */
    /**
     * Sample code: Gets the sensitivity label of a given column in a managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheSensitivityLabelOfAGivenColumnInAManagedDatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSensitivityLabels().getWithResponse("myRG", "myManagedInstanceName",
            "myDatabase", "dbo", "myTable", "myColumn", SensitivityLabelSource.CURRENT,
            com.azure.core.util.Context.NONE);
    }
}
