/** Samples for UnresolvedDependencies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/UnresolvedDependencies_Get.json
     */
    /**
     * Sample code: UnresolvedDependencies_Get.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void unresolvedDependenciesGet(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .unresolvedDependencies()
            .get("rg1", "movecollection1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
