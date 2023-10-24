/** Samples for MoveResources Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Delete.json
     */
    /**
     * Sample code: MoveResources_Delete.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveResourcesDelete(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager.moveResources().delete("rg1", "movecollection1", "moveresourcename1", com.azure.core.util.Context.NONE);
    }
}
