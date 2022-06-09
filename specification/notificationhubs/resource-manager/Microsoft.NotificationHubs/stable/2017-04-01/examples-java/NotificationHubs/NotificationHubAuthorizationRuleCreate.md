```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.notificationhubs.fluent.models.SharedAccessAuthorizationRuleProperties;
import com.azure.resourcemanager.notificationhubs.models.AccessRights;
import com.azure.resourcemanager.notificationhubs.models.SharedAccessAuthorizationRuleCreateOrUpdateParameters;
import java.util.Arrays;

/** Samples for NotificationHubs CreateOrUpdateAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleCreate.json
     */
    /**
     * Sample code: NotificationHubAuthorizationRuleCreate.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubAuthorizationRuleCreate(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .notificationHubs()
            .createOrUpdateAuthorizationRuleWithResponse(
                "5ktrial",
                "nh-sdk-ns",
                "nh-sdk-hub",
                "DefaultListenSharedAccessSignature",
                new SharedAccessAuthorizationRuleCreateOrUpdateParameters()
                    .withProperties(
                        new SharedAccessAuthorizationRuleProperties()
                            .withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-notificationhubs_1.0.0-beta.3/sdk/notificationhubs/azure-resourcemanager-notificationhubs/README.md) on how to add the SDK to your project and authenticate.
