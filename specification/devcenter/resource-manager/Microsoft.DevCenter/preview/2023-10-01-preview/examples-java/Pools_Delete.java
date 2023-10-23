/** Samples for Pools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_Delete.json
     */
    /**
     * Sample code: Pools_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void poolsDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.pools().delete("rg1", "DevProject", "poolName", com.azure.core.util.Context.NONE);
    }
}
