Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-signalr_1.0.0-beta.4/sdk/signalr/azure-resourcemanager-signalr/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for SignalRSharedPrivateLinkResources CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRSharedPrivateLinkResources_CreateOrUpdate.json
     */
    /**
     * Sample code: SignalRSharedPrivateLinkResources_CreateOrUpdate.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRSharedPrivateLinkResourcesCreateOrUpdate(
        com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRSharedPrivateLinkResources()
            .define("upstream")
            .withExistingSignalR("myResourceGroup", "mySignalRService")
            .withGroupId("sites")
            .withPrivateLinkResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Web/sites/myWebApp")
            .withRequestMessage("Please approve")
            .create();
    }
}
```
