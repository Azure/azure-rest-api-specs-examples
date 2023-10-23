/** Samples for Pools ListByProject. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_List.json
     */
    /**
     * Sample code: Pools_ListByProject.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void poolsListByProject(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.pools().listByProject("rg1", "DevProject", null, com.azure.core.util.Context.NONE);
    }
}
