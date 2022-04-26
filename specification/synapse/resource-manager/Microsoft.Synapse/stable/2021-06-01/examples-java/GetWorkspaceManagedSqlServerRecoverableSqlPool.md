Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WorkspaceManagedSqlServerRecoverableSqlPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerRecoverableSqlPool.json
     */
    /**
     * Sample code: Get recoverable sql pools for the server.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getRecoverableSqlPoolsForTheServer(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerRecoverableSqlPools()
            .getWithResponse("wsg-7398", "testWorkspace", "recoverableSqlpools-1235", Context.NONE);
    }
}
```
