Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for OperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/asset-tracks-operation-status-by-id-non-terminal-state.json
     */
    /**
     * Sample code: Get status of asynchronous operation when it is ongoing.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getStatusOfAsynchronousOperationWhenItIsOngoing(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .operationStatuses()
            .getWithResponse(
                "contoso",
                "contosomedia",
                "ClimbingMountRainer",
                "text1",
                "5827d9a1-1fb4-4e54-ac40-8eeed9b862c8",
                Context.NONE);
    }
}
```
