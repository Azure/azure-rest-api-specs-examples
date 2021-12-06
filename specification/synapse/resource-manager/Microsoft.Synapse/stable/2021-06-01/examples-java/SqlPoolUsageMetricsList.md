Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolUsages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolUsageMetricsList.json
     */
    /**
     * Sample code: List the usages of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listTheUsagesOfASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolUsages().list("sqlcrudtest-6730", "sqlcrudtest-9007", "3481", Context.NONE);
    }
}
```
