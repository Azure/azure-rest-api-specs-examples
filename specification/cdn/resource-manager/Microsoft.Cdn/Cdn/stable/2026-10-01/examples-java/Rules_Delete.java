
/**
 * Samples for Rules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Rules_Delete.json
     */
    /**
     * Sample code: Rules_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void rulesDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRules().delete("RG", "profile1", "ruleSet1", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
