import com.azure.core.util.Context;

/** Samples for DigitalTwins List. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/DigitalTwinsList_example.json
     */
    /**
     * Sample code: Get DigitalTwinsInstance resources by subscription.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getDigitalTwinsInstanceResourcesBySubscription(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwins().list(Context.NONE);
    }
}
