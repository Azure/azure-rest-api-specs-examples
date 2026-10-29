
/**
 * Samples for Policies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/WafPolicyDelete.json
     */
    /**
     * Sample code: Delete protection policy.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void deleteProtectionPolicy(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getPolicies().deleteWithResponse("rg1", "Policy1", com.azure.core.util.Context.NONE);
    }
}
