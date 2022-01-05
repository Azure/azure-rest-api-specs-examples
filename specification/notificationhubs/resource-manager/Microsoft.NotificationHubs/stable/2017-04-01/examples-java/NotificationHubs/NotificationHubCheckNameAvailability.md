Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-notificationhubs_1.0.0-beta.3/sdk/notificationhubs/azure-resourcemanager-notificationhubs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.notificationhubs.models.CheckAvailabilityParameters;

/** Samples for NotificationHubs CheckNotificationHubAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubCheckNameAvailability.json
     */
    /**
     * Sample code: notificationHubCheckNameAvailability.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubCheckNameAvailability(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .notificationHubs()
            .checkNotificationHubAvailabilityWithResponse(
                "5ktrial",
                "locp-newns",
                new CheckAvailabilityParameters().withName("sdktest").withLocation("West Europe"),
                Context.NONE);
    }
}
```
