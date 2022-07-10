import com.azure.core.util.Context;

/** Samples for Appliances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesListBySubscription.json
     */
    /**
     * Sample code: List Appliances by subscription.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void listAppliancesBySubscription(
        com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager.appliances().list(Context.NONE);
    }
}
