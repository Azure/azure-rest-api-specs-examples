/** Samples for RuleSets ListResourceUsage. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/RuleSets_ListResourceUsage.json
     */
    /**
     * Sample code: RuleSets_ListResourceUsage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ruleSetsListResourceUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRuleSets()
            .listResourceUsage("RG", "profile1", "ruleSet1", com.azure.core.util.Context.NONE);
    }
}
