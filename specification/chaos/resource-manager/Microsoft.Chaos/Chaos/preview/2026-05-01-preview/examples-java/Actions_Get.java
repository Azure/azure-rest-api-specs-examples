
/**
 * Samples for Actions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Actions_Get.json
     */
    /**
     * Sample code: Get an Action for westus2 location.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAnActionForWestus2Location(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.actions().getWithResponse("westus2", "microsoft-compute-shutdown", com.azure.core.util.Context.NONE);
    }
}
