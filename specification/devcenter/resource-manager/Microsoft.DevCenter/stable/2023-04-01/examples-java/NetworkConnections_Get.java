/** Samples for NetworkConnections GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/NetworkConnections_Get.json
     */
    /**
     * Sample code: NetworkConnections_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .networkConnections()
            .getByResourceGroupWithResponse("rg1", "uswest3network", com.azure.core.util.Context.NONE);
    }
}
