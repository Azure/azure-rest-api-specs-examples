```java
import com.azure.core.util.Context;

/** Samples for WorkspaceManagedIdentitySqlControlSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetManagedIdentitySqlControlSettings.json
     */
    /**
     * Sample code: Get managed identity sql control settings.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getManagedIdentitySqlControlSettings(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedIdentitySqlControlSettings()
            .getWithResponse("resourceGroup1", "workspace1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
