/** Samples for MoveCollections GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Get.json
     */
    /**
     * Sample code: MoveCollections_Get.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsGet(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .moveCollections()
            .getByResourceGroupWithResponse("rg1", "movecollection1", com.azure.core.util.Context.NONE);
    }
}
