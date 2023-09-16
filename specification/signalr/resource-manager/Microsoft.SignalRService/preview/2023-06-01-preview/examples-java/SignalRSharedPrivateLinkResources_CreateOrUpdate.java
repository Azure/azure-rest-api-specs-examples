/** Samples for SignalRSharedPrivateLinkResources CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRSharedPrivateLinkResources_CreateOrUpdate.json
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
