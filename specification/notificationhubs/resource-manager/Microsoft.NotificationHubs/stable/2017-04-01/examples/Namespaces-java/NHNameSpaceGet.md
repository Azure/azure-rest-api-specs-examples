```java
import com.azure.core.util.Context;

/** Samples for Namespaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceGet.json
     */
    /**
     * Sample code: NameSpaceGet.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceGet(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().getByResourceGroupWithResponse("5ktrial", "nh-sdk-ns", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-notificationhubs_1.0.0-beta.2/sdk/notificationhubs/azure-resourcemanager-notificationhubs/README.md) on how to add the SDK to your project and authenticate.
