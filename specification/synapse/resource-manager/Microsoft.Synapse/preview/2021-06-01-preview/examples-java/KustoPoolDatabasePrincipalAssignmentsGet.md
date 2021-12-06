Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for KustoPoolDatabasePrincipalAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasePrincipalAssignmentsGet.json
     */
    /**
     * Sample code: KustoPoolDatabasePrincipalAssignmentsGet.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasePrincipalAssignmentsGet(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabasePrincipalAssignments()
            .getWithResponse(
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "Kustodatabase8",
                "kustoprincipal1",
                "kustorptest",
                Context.NONE);
    }
}
```
