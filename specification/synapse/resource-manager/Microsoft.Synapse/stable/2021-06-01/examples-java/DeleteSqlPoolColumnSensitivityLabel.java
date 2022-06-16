import com.azure.core.util.Context;

/** Samples for SqlPoolSensitivityLabels Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolColumnSensitivityLabel.json
     */
    /**
     * Sample code: Deletes the sensitivity label of a given column.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deletesTheSensitivityLabelOfAGivenColumn(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolSensitivityLabels()
            .deleteWithResponse("myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", Context.NONE);
    }
}
