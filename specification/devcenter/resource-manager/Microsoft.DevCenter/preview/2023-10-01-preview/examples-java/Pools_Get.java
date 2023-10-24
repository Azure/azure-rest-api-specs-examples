/** Samples for Pools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Pools_Get.json
     */
    /**
     * Sample code: Pools_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void poolsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.pools().getWithResponse("rg1", "DevProject", "DevPool", com.azure.core.util.Context.NONE);
    }
}
