Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.TransparentDataEncryptionName;

/** Samples for SqlPoolTransparentDataEncryptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolTransparentDataEncryption.json
     */
    /**
     * Sample code: Get transparent data encryption configuration of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getTransparentDataEncryptionConfigurationOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolTransparentDataEncryptions()
            .getWithResponse(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                TransparentDataEncryptionName.CURRENT,
                Context.NONE);
    }
}
```
