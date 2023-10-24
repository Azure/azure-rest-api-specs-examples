/** Samples for MoveResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_List.json
     */
    /**
     * Sample code: MoveResources_List.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveResourcesList(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager.moveResources().list("rg1", "movecollection1", null, com.azure.core.util.Context.NONE);
    }
}
