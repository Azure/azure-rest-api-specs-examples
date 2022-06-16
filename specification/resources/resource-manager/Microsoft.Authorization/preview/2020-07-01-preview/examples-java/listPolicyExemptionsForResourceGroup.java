import com.azure.core.util.Context;

/** Samples for PolicyExemptions ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForResourceGroup.json
     */
    /**
     * Sample code: List policy exemptions that apply to a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPolicyExemptionsThatApplyToAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .policyClient()
            .getPolicyExemptions()
            .listByResourceGroup("TestResourceGroup", "atScope()", Context.NONE);
    }
}
