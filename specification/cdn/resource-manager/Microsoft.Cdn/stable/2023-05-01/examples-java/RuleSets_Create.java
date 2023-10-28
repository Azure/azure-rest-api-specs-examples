/** Samples for RuleSets Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/RuleSets_Create.json
     */
    /**
     * Sample code: RuleSets_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ruleSetsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRuleSets()
            .createWithResponse("RG", "profile1", "ruleSet1", com.azure.core.util.Context.NONE);
    }
}
