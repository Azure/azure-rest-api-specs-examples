Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.FollowerDatabaseDefinitionInner;

/** Samples for KustoPools DetachFollowerDatabases. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolFollowerDatabasesDetach.json
     */
    /**
     * Sample code: KustoPoolDetachFollowerDatabases.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDetachFollowerDatabases(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .detachFollowerDatabases(
                "kustorptest",
                "kustoclusterrptest4",
                "kustorptest",
                new FollowerDatabaseDefinitionInner()
                    .withKustoPoolResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustoPools/leader4")
                    .withAttachedDatabaseConfigurationName("myAttachedDatabaseConfiguration"),
                Context.NONE);
    }
}
```
