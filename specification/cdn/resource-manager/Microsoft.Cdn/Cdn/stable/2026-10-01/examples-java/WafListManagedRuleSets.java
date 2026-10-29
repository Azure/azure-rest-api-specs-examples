
/**
 * Samples for ManagedRuleSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/WafListManagedRuleSets.json
     */
    /**
     * Sample code: List Policies in a Resource Group.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void listPoliciesInAResourceGroup(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getManagedRuleSets().list(com.azure.core.util.Context.NONE);
    }
}
