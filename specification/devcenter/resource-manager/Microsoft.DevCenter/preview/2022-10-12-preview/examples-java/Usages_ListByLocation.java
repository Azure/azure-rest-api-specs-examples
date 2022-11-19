import com.azure.core.util.Context;

/** Samples for Usages ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/Usages_ListByLocation.json
     */
    /**
     * Sample code: listUsages.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void listUsages(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.usages().listByLocation("westus", Context.NONE);
    }
}
