/** Samples for Pools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Pools_GetUnhealthyStatus.json
     */
    /**
     * Sample code: Pools_GetUnhealthyStatus.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void poolsGetUnhealthyStatus(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.pools().getWithResponse("rg1", "DevProject", "DevPool", com.azure.core.util.Context.NONE);
    }
}
