/** Samples for Pools RunHealthChecks. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_RunHealthChecks.json
     */
    /**
     * Sample code: Pools_RefreshStatus.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void poolsRefreshStatus(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.pools().runHealthChecks("rg1", "DevProject", "DevPool", com.azure.core.util.Context.NONE);
    }
}
