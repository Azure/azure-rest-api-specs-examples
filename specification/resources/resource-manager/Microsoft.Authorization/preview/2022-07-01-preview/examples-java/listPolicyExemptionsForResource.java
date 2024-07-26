
/**
 * Samples for PolicyExemptions ListForResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/
     * listPolicyExemptionsForResource.json
     */
    /**
     * Sample code: List all policy exemptions that apply to a resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllPolicyExemptionsThatApplyToAResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyExemptions().listForResource("TestResourceGroup",
            "Microsoft.Compute", "virtualMachines/MyTestVm", "domainNames", "MyTestComputer.cloudapp.net", null,
            com.azure.core.util.Context.NONE);
    }
}
