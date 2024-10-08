
import com.azure.resourcemanager.sql.models.RecommendedSensitivityLabelUpdate;
import com.azure.resourcemanager.sql.models.RecommendedSensitivityLabelUpdateKind;
import com.azure.resourcemanager.sql.models.RecommendedSensitivityLabelUpdateList;
import java.util.Arrays;

/**
 * Samples for ManagedDatabaseRecommendedSensitivityLabels Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseSensitivityLabelsRecommendedUpdate.json
     */
    /**
     * Sample code: Update recommended sensitivity labels of a given database using an operations batch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateRecommendedSensitivityLabelsOfAGivenDatabaseUsingAnOperationsBatch(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseRecommendedSensitivityLabels()
            .updateWithResponse("myRG", "myManagedInstanceName", "myDatabase",
                new RecommendedSensitivityLabelUpdateList().withOperations(Arrays.asList(
                    new RecommendedSensitivityLabelUpdate().withOp(RecommendedSensitivityLabelUpdateKind.ENABLE)
                        .withSchema("dbo").withTable("table1").withColumn("column1"),
                    new RecommendedSensitivityLabelUpdate().withOp(RecommendedSensitivityLabelUpdateKind.DISABLE)
                        .withSchema("dbo").withTable("table2").withColumn("column2"),
                    new RecommendedSensitivityLabelUpdate().withOp(RecommendedSensitivityLabelUpdateKind.DISABLE)
                        .withSchema("dbo").withTable("Table1").withColumn("Column3"))),
                com.azure.core.util.Context.NONE);
    }
}
