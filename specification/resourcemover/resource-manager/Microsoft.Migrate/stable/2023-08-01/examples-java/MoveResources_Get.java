/** Samples for MoveResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Get.json
     */
    /**
     * Sample code: MoveResources_Get.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveResourcesGet(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .moveResources()
            .getWithResponse("rg1", "movecollection1", "moveresourcename1", com.azure.core.util.Context.NONE);
    }
}
