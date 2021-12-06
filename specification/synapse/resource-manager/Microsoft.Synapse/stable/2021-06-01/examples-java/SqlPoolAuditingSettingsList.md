Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolBlobAuditingPolicies ListBySqlPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolAuditingSettingsList.json
     */
    /**
     * Sample code: List audit settings of a database.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listAuditSettingsOfADatabase(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolBlobAuditingPolicies()
            .listBySqlPool("blobauditingtest-6852", "blobauditingtest-2080", "testdb", Context.NONE);
    }
}
```
