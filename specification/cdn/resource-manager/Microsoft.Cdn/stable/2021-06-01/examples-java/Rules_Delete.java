import com.azure.core.util.Context;

/** Samples for Rules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Delete.json
     */
    /**
     * Sample code: Rules_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rulesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRules()
            .delete("RG", "profile1", "ruleSet1", "rule1", Context.NONE);
    }
}
