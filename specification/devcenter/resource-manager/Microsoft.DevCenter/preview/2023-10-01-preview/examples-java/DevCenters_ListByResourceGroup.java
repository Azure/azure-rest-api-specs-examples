/** Samples for DevCenters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/DevCenters_ListByResourceGroup.json
     */
    /**
     * Sample code: DevCenters_ListByResourceGroup.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devCentersListByResourceGroup(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.devCenters().listByResourceGroup("rg1", null, com.azure.core.util.Context.NONE);
    }
}
