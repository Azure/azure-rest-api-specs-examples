/** Samples for MoveCollections ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ListMoveCollectionsByResourceGroup.json
     */
    /**
     * Sample code: MoveCollections_ListMoveCollectionsByResourceGroup.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsListMoveCollectionsByResourceGroup(
        com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager.moveCollections().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
