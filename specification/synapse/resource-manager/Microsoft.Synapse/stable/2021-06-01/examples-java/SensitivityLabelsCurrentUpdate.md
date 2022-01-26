Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.SensitivityLabelInner;
import com.azure.resourcemanager.synapse.fluent.models.SensitivityLabelUpdateInner;
import com.azure.resourcemanager.synapse.models.SensitivityLabelRank;
import com.azure.resourcemanager.synapse.models.SensitivityLabelUpdateKind;
import com.azure.resourcemanager.synapse.models.SensitivityLabelUpdateList;
import java.util.Arrays;

/** Samples for SqlPoolSensitivityLabels Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsCurrentUpdate.json
     */
    /**
     * Sample code: Update sensitivity labels of a given database using an operations batch.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateSensitivityLabelsOfAGivenDatabaseUsingAnOperationsBatch(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolSensitivityLabels()
            .updateWithResponse(
                "myRG",
                "myWorkspace",
                "mySqlPool",
                new SensitivityLabelUpdateList()
                    .withOperations(
                        Arrays
                            .asList(
                                new SensitivityLabelUpdateInner()
                                    .withOp(SensitivityLabelUpdateKind.SET)
                                    .withSchema("dbo")
                                    .withTable("table1")
                                    .withColumn("column1")
                                    .withSensitivityLabel(
                                        new SensitivityLabelInner()
                                            .withLabelName("Highly Confidential")
                                            .withLabelId("3A477B16-9423-432B-AA97-6069B481CEC3")
                                            .withInformationType("Financial")
                                            .withInformationTypeId("1D3652D6-422C-4115-82F1-65DAEBC665C8")
                                            .withRank(SensitivityLabelRank.LOW)),
                                new SensitivityLabelUpdateInner()
                                    .withOp(SensitivityLabelUpdateKind.SET)
                                    .withSchema("dbo")
                                    .withTable("table2")
                                    .withColumn("column2")
                                    .withSensitivityLabel(
                                        new SensitivityLabelInner()
                                            .withLabelName("PII")
                                            .withLabelId("bf91e08c-f4f0-478a-b016-25164b2a65ff")
                                            .withInformationType("PhoneNumber")
                                            .withInformationTypeId("d22fa6e9-5ee4-3bde-4c2b-a409604c4646")
                                            .withRank(SensitivityLabelRank.CRITICAL)),
                                new SensitivityLabelUpdateInner()
                                    .withOp(SensitivityLabelUpdateKind.REMOVE)
                                    .withSchema("dbo")
                                    .withTable("Table1")
                                    .withColumn("Column3"))),
                Context.NONE);
    }
}
```
