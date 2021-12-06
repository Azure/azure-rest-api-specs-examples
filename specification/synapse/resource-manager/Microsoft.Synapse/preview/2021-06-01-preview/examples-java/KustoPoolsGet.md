Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for KustoPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsGet.json
     */
    /**
     * Sample code: kustoPoolsGet.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsGet(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .getWithResponse("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", Context.NONE);
    }
}
```
