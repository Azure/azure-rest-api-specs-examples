/** Samples for DevCenters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/DevCenters_Delete.json
     */
    /**
     * Sample code: DevCenters_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devCentersDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.devCenters().delete("rg1", "Contoso", com.azure.core.util.Context.NONE);
    }
}
