
import com.azure.resourcemanager.sql.models.CurrentSensitivityLabelSource;

/**
 * Samples for SensitivityLabels Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ColumnSensitivityLabelDelete.json
     */
    /**
     * Sample code: Deletes the sensitivity label of a given column.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        deletesTheSensitivityLabelOfAGivenColumn(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSensitivityLabels().deleteWithResponse("myRG", "myServer", "myDatabase", "dbo",
            "myTable", "myColumn", CurrentSensitivityLabelSource.CURRENT, com.azure.core.util.Context.NONE);
    }
}
