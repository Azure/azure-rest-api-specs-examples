
/**
 * Samples for Policies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/WafListPolicies.json
     */
    /**
     * Sample code: List Policies in a Resource Group.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void listPoliciesInAResourceGroup(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getPolicies().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
