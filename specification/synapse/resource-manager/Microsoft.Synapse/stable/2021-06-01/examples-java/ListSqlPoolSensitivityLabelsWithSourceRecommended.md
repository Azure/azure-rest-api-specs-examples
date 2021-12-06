Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolSensitivityLabels ListRecommended. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolSensitivityLabelsWithSourceRecommended.json
     */
    /**
     * Sample code: Gets the recommended sensitivity labels of a given SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsTheRecommendedSensitivityLabelsOfAGivenSQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolSensitivityLabels()
            .listRecommended("myRG", "myServer", "myDatabase", null, null, null, Context.NONE);
    }
}
```
