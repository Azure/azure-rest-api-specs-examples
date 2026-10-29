
/**
 * Samples for Rules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Rules_Get.json
     */
    /**
     * Sample code: Rules_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void rulesGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRules().getWithResponse("RG", "profile1", "ruleSet1", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
