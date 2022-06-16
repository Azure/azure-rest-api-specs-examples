import com.azure.core.util.Context;

/** Samples for RuleSets ListByProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/RuleSets_ListByProfile.json
     */
    /**
     * Sample code: RuleSets_ListByProfile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ruleSetsListByProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getRuleSets().listByProfile("RG", "profile1", Context.NONE);
    }
}
