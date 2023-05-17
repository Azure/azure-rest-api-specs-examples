/** Samples for DevCenters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/DevCenters_ListBySubscription.json
     */
    /**
     * Sample code: DevCenters_ListBySubscription.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devCentersListBySubscription(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.devCenters().list(null, com.azure.core.util.Context.NONE);
    }
}
