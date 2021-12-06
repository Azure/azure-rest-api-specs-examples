Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-notificationhubs_1.0.0-beta.2/sdk/notificationhubs/azure-resourcemanager-notificationhubs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for NotificationHubs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubGet.json
     */
    /**
     * Sample code: NotificationHubGet.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubGet(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.notificationHubs().getWithResponse("5ktrial", "nh-sdk-ns", "nh-sdk-hub", Context.NONE);
    }
}
```
