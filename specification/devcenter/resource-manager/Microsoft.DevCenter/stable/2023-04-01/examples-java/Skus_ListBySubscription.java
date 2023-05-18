/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Skus_ListBySubscription.json
     */
    /**
     * Sample code: Skus_ListBySubscription.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void skusListBySubscription(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.skus().list(null, com.azure.core.util.Context.NONE);
    }
}
