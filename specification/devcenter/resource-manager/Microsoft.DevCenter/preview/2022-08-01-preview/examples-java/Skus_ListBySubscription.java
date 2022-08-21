import com.azure.core.util.Context;

/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Skus_ListBySubscription.json
     */
    /**
     * Sample code: Skus_ListBySubscription.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void skusListBySubscription(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.skus().list(null, Context.NONE);
    }
}
