Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-signalr_1.0.0-beta.3/sdk/signalr/azure-resourcemanager-signalr/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SignalR ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/SignalR_ListSkus.json
     */
    /**
     * Sample code: SignalR_ListSkus.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListSkus(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().listSkusWithResponse("myResourceGroup", "mySignalRService", Context.NONE);
    }
}
```
