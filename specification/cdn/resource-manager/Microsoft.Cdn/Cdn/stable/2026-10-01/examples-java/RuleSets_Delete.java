
/**
 * Samples for RuleSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RuleSets_Delete.json
     */
    /**
     * Sample code: RuleSets_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void ruleSetsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRuleSets().delete("RG", "profile1", "ruleSet1", com.azure.core.util.Context.NONE);
    }
}
