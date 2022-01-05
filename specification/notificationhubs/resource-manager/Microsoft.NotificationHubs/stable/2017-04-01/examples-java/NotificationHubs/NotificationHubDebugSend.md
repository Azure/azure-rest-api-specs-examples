Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-notificationhubs_1.0.0-beta.3/sdk/notificationhubs/azure-resourcemanager-notificationhubs/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/** Samples for NotificationHubs DebugSend. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubDebugSend.json
     */
    /**
     * Sample code: debugsend.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void debugsend(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager)
        throws IOException {
        manager
            .notificationHubs()
            .debugSendWithResponse(
                "5ktrial",
                "nh-sdk-ns",
                "nh-sdk-hub",
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize("{\"data\":{\"message\":\"Hello\"}}", Object.class, SerializerEncoding.JSON),
                Context.NONE);
    }
}
```
