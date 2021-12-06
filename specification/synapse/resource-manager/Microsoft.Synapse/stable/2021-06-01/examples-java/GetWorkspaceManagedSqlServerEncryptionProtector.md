Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.EncryptionProtectorName;

/** Samples for WorkspaceManagedSqlServerEncryptionProtector Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerEncryptionProtector.json
     */
    /**
     * Sample code: Get workspace managed sql Server's encryption protector.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getWorkspaceManagedSqlServerSEncryptionProtector(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerEncryptionProtectors()
            .getWithResponse("wsg-7398", "testWorkspace", EncryptionProtectorName.CURRENT, Context.NONE);
    }
}
```
