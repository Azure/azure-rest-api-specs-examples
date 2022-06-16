import com.azure.core.util.Context;

/** Samples for ManagedRuleSets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafListManagedRuleSets.json
     */
    /**
     * Sample code: List Policies in a Resource Group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPoliciesInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getManagedRuleSets().list(Context.NONE);
    }
}
