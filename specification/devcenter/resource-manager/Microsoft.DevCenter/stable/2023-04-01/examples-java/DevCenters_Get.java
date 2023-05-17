/** Samples for DevCenters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/DevCenters_Get.json
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
