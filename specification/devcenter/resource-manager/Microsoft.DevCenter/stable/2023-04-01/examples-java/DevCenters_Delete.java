/** Samples for DevCenters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/DevCenters_Delete.json
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
