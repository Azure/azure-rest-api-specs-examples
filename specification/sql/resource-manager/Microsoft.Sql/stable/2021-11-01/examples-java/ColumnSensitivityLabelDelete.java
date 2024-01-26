
import com.azure.core.util.Context;

/** Samples for SensitivityLabels Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ColumnSensitivityLabelDelete.json
     */
    /**
     * Sample code: Deletes the sensitivity label of a given column.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesTheSensitivityLabelOfAGivenColumn(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSensitivityLabels().deleteWithResponse("myRG", "myServer",
            "myDatabase", "dbo", "myTable", "myColumn", Context.NONE);
    }
}
