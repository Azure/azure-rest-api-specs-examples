
/**
 * Samples for RuleSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RuleSets_Get.json
     */
    /**
     * Sample code: RuleSets_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void ruleSetsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRuleSets().getWithResponse("RG", "profile1", "ruleSet1",
            com.azure.core.util.Context.NONE);
    }
}
