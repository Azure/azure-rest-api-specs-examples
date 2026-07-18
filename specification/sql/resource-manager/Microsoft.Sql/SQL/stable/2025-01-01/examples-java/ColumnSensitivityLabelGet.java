
import com.azure.resourcemanager.sql.models.SensitivityLabelSource;

/**
 * Samples for SensitivityLabels Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ColumnSensitivityLabelGet.json
     */
    /**
     * Sample code: Gets the sensitivity label of a given column.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheSensitivityLabelOfAGivenColumn(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSensitivityLabels().getWithResponse("myRG", "myServer", "myDatabase", "dbo",
            "myTable", "myColumn", SensitivityLabelSource.CURRENT, com.azure.core.util.Context.NONE);
    }
}
