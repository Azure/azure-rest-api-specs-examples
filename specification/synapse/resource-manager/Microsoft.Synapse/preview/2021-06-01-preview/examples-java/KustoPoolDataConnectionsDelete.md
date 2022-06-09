```java
import com.azure.core.util.Context;

/** Samples for KustoPoolDataConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsDelete.json
     */
    /**
     * Sample code: KustoPoolDataConnectionsDelete.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDataConnectionsDelete(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDataConnections()
            .delete(
                "kustorptest",
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "KustoDatabase8",
                "kustoeventhubconnection1",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
