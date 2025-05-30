
/**
 * Samples for TargetTypes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/TargetTypes_Get.json
     */
    /**
     * Sample code: Get a Target Type for westus2 location.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getATargetTypeForWestus2Location(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.targetTypes().getWithResponse("westus2", "Microsoft-Agent", com.azure.core.util.Context.NONE);
    }
}
