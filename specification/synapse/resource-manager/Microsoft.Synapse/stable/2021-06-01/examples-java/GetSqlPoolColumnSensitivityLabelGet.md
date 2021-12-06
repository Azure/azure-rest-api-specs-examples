Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.SensitivityLabelSource;

/** Samples for SqlPoolSensitivityLabels Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolColumnSensitivityLabelGet.json
     */
    /**
     * Sample code: Gets the sensitivity label of a given column.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsTheSensitivityLabelOfAGivenColumn(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolSensitivityLabels()
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
```
