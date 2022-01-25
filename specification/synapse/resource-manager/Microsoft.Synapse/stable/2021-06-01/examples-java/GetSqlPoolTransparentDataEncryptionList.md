Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolTransparentDataEncryptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolTransparentDataEncryptionList.json
     */
    /**
     * Sample code: Get list of transparent data encryption configurations of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getListOfTransparentDataEncryptionConfigurationsOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolTransparentDataEncryptions()
            .list("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", Context.NONE);
    }
}
```
