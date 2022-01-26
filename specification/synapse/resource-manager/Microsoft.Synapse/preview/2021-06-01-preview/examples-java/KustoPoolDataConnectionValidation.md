Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.DataConnectionValidationInner;
import com.azure.resourcemanager.synapse.models.EventHubDataConnection;

/** Samples for KustoPoolDataConnections DataConnectionValidation. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionValidation.json
     */
    /**
     * Sample code: KustoPoolDataConnectionValidation.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDataConnectionValidation(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDataConnections()
            .dataConnectionValidation(
                "kustorptest",
                "kustorptest",
                "kustoclusterrptest4",
                "KustoDatabase8",
                new DataConnectionValidationInner()
                    .withDataConnectionName("DataConnections8")
                    .withProperties(new EventHubDataConnection()),
                Context.NONE);
    }
}
```
