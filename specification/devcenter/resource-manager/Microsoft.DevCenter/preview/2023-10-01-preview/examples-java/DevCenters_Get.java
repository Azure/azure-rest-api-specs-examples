/** Samples for DevCenters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/DevCenters_Get.json
     */
    /**
     * Sample code: DevCenters_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devCentersGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.devCenters().getByResourceGroupWithResponse("rg1", "Contoso", com.azure.core.util.Context.NONE);
    }
}
