
import com.azure.resourcemanager.sql.fluent.models.SensitivityLabelInner;
import com.azure.resourcemanager.sql.fluent.models.SensitivityLabelUpdateInner;
import com.azure.resourcemanager.sql.models.SensitivityLabelUpdateKind;
import com.azure.resourcemanager.sql.models.SensitivityLabelUpdateList;
import java.util.Arrays;

/**
 * Samples for ManagedDatabaseSensitivityLabels Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseSensitivityLabelsCurrentUpdate.json
     */
    /**
     * Sample code: Update sensitivity labels of a given database using an operations batch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateSensitivityLabelsOfAGivenDatabaseUsingAnOperationsBatch(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSensitivityLabels().updateWithResponse("myRG",
            "myManagedInstanceName", "myDatabase",
            new SensitivityLabelUpdateList().withOperations(Arrays.asList(
                new SensitivityLabelUpdateInner().withOp(SensitivityLabelUpdateKind.SET).withSchema("dbo")
                    .withTable("table1").withColumn("column1")
                    .withSensitivityLabel(new SensitivityLabelInner().withLabelName("Highly Confidential")
                        .withLabelId("3A477B16-9423-432B-AA97-6069B481CEC3").withInformationType("Financial")
                        .withInformationTypeId("1D3652D6-422C-4115-82F1-65DAEBC665C8")),
                new SensitivityLabelUpdateInner().withOp(SensitivityLabelUpdateKind.SET).withSchema("dbo")
                    .withTable("table2").withColumn("column2")
                    .withSensitivityLabel(new SensitivityLabelInner().withLabelName("PII")
                        .withLabelId("bf91e08c-f4f0-478a-b016-25164b2a65ff").withInformationType("PhoneNumber")
                        .withInformationTypeId("d22fa6e9-5ee4-3bde-4c2b-a409604c4646")),
                new SensitivityLabelUpdateInner().withOp(SensitivityLabelUpdateKind.REMOVE).withSchema("dbo")
                    .withTable("Table1").withColumn("Column3"))),
            com.azure.core.util.Context.NONE);
    }
}
