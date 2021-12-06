Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolOperationResults GetLocationHeaderResult. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetLocationHeaderResultWithSqlPool.json
     */
    /**
     * Sample code: Get the result of an operation on a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getTheResultOfAnOperationOnASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolOperationResults()
            .getLocationHeaderResultWithResponse(
                "ExampleResourceGroup",
                "ExampleWorkspace",
                "ExampleSqlPool",
                "fedcba98-7654-4210-fedc-ba9876543210",
                Context.NONE);
    }
}
```
