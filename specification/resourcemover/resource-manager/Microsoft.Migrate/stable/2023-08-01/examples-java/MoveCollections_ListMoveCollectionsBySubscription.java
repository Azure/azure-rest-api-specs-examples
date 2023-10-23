/** Samples for MoveCollections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ListMoveCollectionsBySubscription.json
     */
    /**
     * Sample code: MoveCollections_ListMoveCollectionsBySubscription.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsListMoveCollectionsBySubscription(
        com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager.moveCollections().list(com.azure.core.util.Context.NONE);
    }
}
