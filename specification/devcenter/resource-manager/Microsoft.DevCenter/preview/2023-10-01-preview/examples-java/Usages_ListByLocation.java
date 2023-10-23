/** Samples for Usages ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Usages_ListByLocation.json
     */
    /**
     * Sample code: listUsages.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void listUsages(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.usages().listByLocation("westus", com.azure.core.util.Context.NONE);
    }
}
