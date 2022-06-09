```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.batch.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.batch.models.PrivateLinkServiceConnectionStatus;

/** Samples for PrivateEndpointConnection Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PrivateEndpointConnectionUpdate.json
     */
    /**
     * Sample code: UpdatePrivateEndpointConnection.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void updatePrivateEndpointConnection(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .privateEndpointConnections()
            .update(
                "default-azurebatch-japaneast",
                "sampleacct",
                "testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0",
                new PrivateEndpointConnectionInner()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                            .withDescription("Approved by xyz.abc@company.com")),
                null,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
