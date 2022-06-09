```java
/** Samples for NotificationHubs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubCreate.json
     */
    /**
     * Sample code: NotificationHubCreate.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void notificationHubCreate(
        com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .notificationHubs()
            .define("nh-sdk-hub")
            .withRegion("eastus")
            .withExistingNamespace("5ktrial", "nh-sdk-ns")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-notificationhubs_1.0.0-beta.3/sdk/notificationhubs/azure-resourcemanager-notificationhubs/README.md) on how to add the SDK to your project and authenticate.
