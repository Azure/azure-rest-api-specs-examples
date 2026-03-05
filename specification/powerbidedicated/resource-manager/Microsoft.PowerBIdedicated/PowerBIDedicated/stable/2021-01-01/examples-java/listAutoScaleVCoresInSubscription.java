
/**
 * Samples for AutoScaleVCores List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-01-01/listAutoScaleVCoresInSubscription.json
     */
    /**
     * Sample code: List auto scale v-cores in subscription.
     * 
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void
        listAutoScaleVCoresInSubscription(com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.autoScaleVCores().list(com.azure.core.util.Context.NONE);
    }
}
