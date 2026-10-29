
/**
 * Samples for Policies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/WafPolicyGet.json
     */
    /**
     * Sample code: Get Policy.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void getPolicy(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getPolicies().getByResourceGroupWithResponse("rg1", "MicrosoftCdnWafPolicy",
            com.azure.core.util.Context.NONE);
    }
}
