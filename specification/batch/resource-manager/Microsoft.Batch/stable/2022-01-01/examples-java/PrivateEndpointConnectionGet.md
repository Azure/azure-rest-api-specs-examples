Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnection Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PrivateEndpointConnectionGet.json
     */
    /**
     * Sample code: GetPrivateEndpointConnection.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void getPrivateEndpointConnection(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse(
                "default-azurebatch-japaneast",
                "sampleacct",
                "testprivateEndpointConnection5testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0",
                Context.NONE);
    }
}
```
