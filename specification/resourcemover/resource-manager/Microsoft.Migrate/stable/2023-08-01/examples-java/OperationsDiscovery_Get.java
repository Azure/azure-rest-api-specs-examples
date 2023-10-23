/** Samples for OperationsDiscovery Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/OperationsDiscovery_Get.json
     */
    /**
     * Sample code: OperationsDiscovery_Get.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void operationsDiscoveryGet(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager.operationsDiscoveries().getWithResponse(com.azure.core.util.Context.NONE);
    }
}
