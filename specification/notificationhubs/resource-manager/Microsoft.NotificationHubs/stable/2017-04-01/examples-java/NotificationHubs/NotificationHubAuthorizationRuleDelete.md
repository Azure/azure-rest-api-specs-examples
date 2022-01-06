Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-notificationhubs_1.0.0-beta.3/sdk/notificationhubs/azure-resourcemanager-notificationhubs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for NotificationHubs DeleteAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleDelete.json
     */
    /**
     * Sample code: NotificationHubAuthorizationRuleDelete.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubAuthorizationRuleDelete(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .notificationHubs()
            .deleteAuthorizationRuleWithResponse(
                "5ktrial", "nh-sdk-ns", "nh-sdk-hub", "DefaultListenSharedAccessSignature", Context.NONE);
    }
}
```
