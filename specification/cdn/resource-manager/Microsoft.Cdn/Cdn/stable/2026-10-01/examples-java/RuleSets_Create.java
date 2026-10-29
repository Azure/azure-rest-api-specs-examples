
/**
 * Samples for RuleSets Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RuleSets_Create.json
     */
    /**
     * Sample code: RuleSets_Create.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void ruleSetsCreate(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRuleSets().create("RG", "profile1", "ruleSet1", null,
            com.azure.core.util.Context.NONE);
    }
}
