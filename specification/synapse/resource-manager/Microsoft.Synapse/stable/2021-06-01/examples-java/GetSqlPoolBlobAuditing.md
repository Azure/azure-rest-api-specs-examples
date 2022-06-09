```java
import com.azure.core.util.Context;

/** Samples for SqlPoolBlobAuditingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolBlobAuditing.json
     */
    /**
     * Sample code: Get blob auditing policy of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getBlobAuditingPolicyOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolBlobAuditingPolicies()
            .getWithResponse("blobauditingtest-6852", "blobauditingtest-2080", "testdb", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
