
/**
 * Samples for RuleSets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/RuleSets_Delete.json
     */
    /**
     * Sample code: RuleSets_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ruleSetsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getRuleSets().delete("RG", "profile1", "ruleSet1",
            com.azure.core.util.Context.NONE);
    }
}
