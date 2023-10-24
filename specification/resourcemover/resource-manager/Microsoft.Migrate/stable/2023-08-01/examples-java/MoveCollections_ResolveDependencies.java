/** Samples for MoveCollections ResolveDependencies. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ResolveDependencies.json
     */
    /**
     * Sample code: MoveCollections_ResolveDependencies.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsResolveDependencies(
        com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager.moveCollections().resolveDependencies("rg1", "movecollection1", com.azure.core.util.Context.NONE);
    }
}
