/** Samples for Pools ListByProject. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Pools_List.json
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
