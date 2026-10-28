
/**
 * Samples for Rules ListByRuleSet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Rules_ListByRuleSet.json
     */
    /**
     * Sample code: Rules_ListByRuleSet.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void rulesListByRuleSet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRules().listByRuleSet("RG", "profile1", "ruleSet1",
            com.azure.core.util.Context.NONE);
    }
}
