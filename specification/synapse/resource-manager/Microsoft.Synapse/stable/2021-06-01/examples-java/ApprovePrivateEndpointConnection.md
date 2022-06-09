```java
import com.azure.resourcemanager.synapse.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ApprovePrivateEndpointConnection.json
     */
    /**
     * Sample code: Approve private endpoint connection.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void approvePrivateEndpointConnection(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateEndpointConnections()
            .define("ExamplePrivateEndpointConnection")
            .withExistingWorkspace("ExampleResourceGroup", "ExampleWorkspace")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus("Approved")
                    .withDescription("Approved by abc@example.com"))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
