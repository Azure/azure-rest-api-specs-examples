Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.synapse.models.TransparentDataEncryptionName;
import com.azure.resourcemanager.synapse.models.TransparentDataEncryptionStatus;

/** Samples for SqlPoolTransparentDataEncryptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolTransparentDataEncryption.json
     */
    /**
     * Sample code: Create or update a Sql pool's transparent data encryption configuration.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateASqlPoolSTransparentDataEncryptionConfiguration(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolTransparentDataEncryptions()
            .define(TransparentDataEncryptionName.CURRENT)
            .withExistingSqlPool("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187")
            .withStatus(TransparentDataEncryptionStatus.ENABLED)
            .create();
    }
}
```
