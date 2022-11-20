import com.azure.core.util.Context;

/** Samples for Pools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/Pools_Get.json
     */
    /**
     * Sample code: Pools_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void poolsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.pools().getWithResponse("rg1", "{projectName}", "{poolName}", Context.NONE);
    }
}
