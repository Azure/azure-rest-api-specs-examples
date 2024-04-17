
/**
 * Samples for RuleSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/RuleSets_Get.json
     */
    /**
     * Sample code: RuleSets_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ruleSetsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getRuleSets().getWithResponse("RG", "profile1", "ruleSet1",
            com.azure.core.util.Context.NONE);
    }
}
