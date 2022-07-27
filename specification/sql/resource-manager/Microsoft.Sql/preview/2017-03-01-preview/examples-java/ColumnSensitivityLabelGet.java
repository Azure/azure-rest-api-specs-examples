import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.SensitivityLabelSource;

/** Samples for SensitivityLabels Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ColumnSensitivityLabelGet.json
     */
    /**
     * Sample code: Gets the sensitivity label of a given column.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheSensitivityLabelOfAGivenColumn(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getSensitivityLabels()
            .getWithResponse(
                "myRG",
                "myServer",
                "myDatabase",
                "dbo",
                "myTable",
                "myColumn",
                SensitivityLabelSource.CURRENT,
                Context.NONE);
    }
}
