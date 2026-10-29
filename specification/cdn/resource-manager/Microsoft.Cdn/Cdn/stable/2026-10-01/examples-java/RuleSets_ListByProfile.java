
/**
 * Samples for RuleSets ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/RuleSets_ListByProfile.json
     */
    /**
     * Sample code: RuleSets_ListByProfile.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void ruleSetsListByProfile(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRuleSets().listByProfile("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
