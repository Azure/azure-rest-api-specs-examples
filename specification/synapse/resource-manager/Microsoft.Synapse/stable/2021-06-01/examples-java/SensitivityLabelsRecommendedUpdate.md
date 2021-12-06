Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.RecommendedSensitivityLabelUpdate;
import com.azure.resourcemanager.synapse.models.RecommendedSensitivityLabelUpdateKind;
import com.azure.resourcemanager.synapse.models.RecommendedSensitivityLabelUpdateList;
import java.util.Arrays;

/** Samples for SqlPoolRecommendedSensitivityLabels Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsRecommendedUpdate.json
     */
    /**
     * Sample code: Update recommended sensitivity labels of a given SQL Pool using an operations batch.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateRecommendedSensitivityLabelsOfAGivenSQLPoolUsingAnOperationsBatch(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolRecommendedSensitivityLabels()
            .updateWithResponse(
                "myRG",
                "myWorkspace",
                "mySqlPool",
                new RecommendedSensitivityLabelUpdateList()
                    .withOperations(
                        Arrays
                            .asList(
                                new RecommendedSensitivityLabelUpdate()
                                    .withOp(RecommendedSensitivityLabelUpdateKind.ENABLE)
                                    .withSchema("dbo")
                                    .withTable("table1")
                                    .withColumn("column1"),
                                new RecommendedSensitivityLabelUpdate()
                                    .withOp(RecommendedSensitivityLabelUpdateKind.ENABLE)
                                    .withSchema("dbo")
                                    .withTable("table2")
                                    .withColumn("column2"),
                                new RecommendedSensitivityLabelUpdate()
                                    .withOp(RecommendedSensitivityLabelUpdateKind.DISABLE)
                                    .withSchema("dbo")
                                    .withTable("table1")
                                    .withColumn("column3"))),
                Context.NONE);
    }
}
```
