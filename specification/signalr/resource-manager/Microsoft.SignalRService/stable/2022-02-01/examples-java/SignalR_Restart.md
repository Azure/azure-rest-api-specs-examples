```java
import com.azure.core.util.Context;

/** Samples for SignalR Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_Restart.json
     */
    /**
     * Sample code: SignalR_Restart.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRRestart(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().restart("myResourceGroup", "mySignalRService", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-signalr_1.0.0-beta.4/sdk/signalr/azure-resourcemanager-signalr/README.md) on how to add the SDK to your project and authenticate.
