
/**
 * Samples for RuleSets ListResourceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RuleSets_ListResourceUsage.json
     */
    /**
     * Sample code: RuleSets_ListResourceUsage.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void ruleSetsListResourceUsage(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRuleSets().listResourceUsage("RG", "profile1", "ruleSet1",
            com.azure.core.util.Context.NONE);
    }
}
