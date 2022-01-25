Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.ConnectionPolicyName;

/** Samples for SqlPoolConnectionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolConnectionPolicy.json
     */
    /**
     * Sample code: Get a connection policy of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAConnectionPolicyOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolConnectionPolicies()
            .getWithResponse(
                "blobauditingtest-6852", "blobauditingtest-2080", "testdb", ConnectionPolicyName.DEFAULT, Context.NONE);
    }
}
```
