/** Samples for MoveCollections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Delete.json
     */
    /**
     * Sample code: MoveCollections_Delete.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsDelete(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager.moveCollections().delete("rg1", "movecollection1", com.azure.core.util.Context.NONE);
    }
}
