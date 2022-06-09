```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.notificationhubs.models.NotificationHubResource;

/** Samples for NotificationHubs Patch. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubPatch.json
     */
    /**
     * Sample code: NotificationHubPatch.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubPatch(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        NotificationHubResource resource =
            manager
                .notificationHubs()
                .getWithResponse("sdkresourceGroup", "nh-sdk-ns", "sdk-notificationHubs-8708", Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-notificationhubs_1.0.0-beta.3/sdk/notificationhubs/azure-resourcemanager-notificationhubs/README.md) on how to add the SDK to your project and authenticate.
