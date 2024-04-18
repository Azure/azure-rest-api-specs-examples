
/**
 * Samples for Policies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/WafPolicyGet.json
     */
    /**
     * Sample code: Get Policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getPolicies().getByResourceGroupWithResponse("rg1",
            "MicrosoftCdnWafPolicy", com.azure.core.util.Context.NONE);
    }
}
