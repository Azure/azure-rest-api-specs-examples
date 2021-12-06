Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.EncryptionProtector;
import com.azure.resourcemanager.synapse.models.EncryptionProtectorName;
import com.azure.resourcemanager.synapse.models.ServerKeyType;

/** Samples for WorkspaceManagedSqlServerEncryptionProtector CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateKeyVault.json
     */
    /**
     * Sample code: Update the encryption protector to key vault.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateTheEncryptionProtectorToKeyVault(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        EncryptionProtector resource =
            manager
                .workspaceManagedSqlServerEncryptionProtectors()
                .getWithResponse("wsg-7398", "testWorkspace", EncryptionProtectorName.CURRENT, Context.NONE)
                .getValue();
        resource
            .update()
            .withServerKeyName("someVault_someKey_01234567890123456789012345678901")
            .withServerKeyType(ServerKeyType.AZURE_KEY_VAULT)
            .apply();
    }
}
```
