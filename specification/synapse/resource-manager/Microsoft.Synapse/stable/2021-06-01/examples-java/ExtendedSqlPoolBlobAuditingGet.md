Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ExtendedSqlPoolBlobAuditingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ExtendedSqlPoolBlobAuditingGet.json
     */
    /**
     * Sample code: Get an extended database's blob auditing policy.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAnExtendedDatabaseSBlobAuditingPolicy(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .extendedSqlPoolBlobAuditingPolicies()
            .getWithResponse("blobauditingtest-6852", "blobauditingtest-2080", "testdb", Context.NONE);
    }
}
```
