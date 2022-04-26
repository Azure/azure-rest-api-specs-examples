Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolSensitivityLabels ListCurrent. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolsSensitivityLabelsWithSourceCurrent.json
     */
    /**
     * Sample code: Gets the current sensitivity labels of a given SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsTheCurrentSensitivityLabelsOfAGivenSQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolSensitivityLabels().listCurrent("myRG", "myServer", "myDatabase", null, Context.NONE);
    }
}
```
