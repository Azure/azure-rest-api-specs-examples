
/**
 * Samples for PolicyAssignments ListForResource.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/
     * listPolicyAssignmentsForResource.json
     */
    /**
     * Sample code: List all policy assignments that apply to a resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllPolicyAssignmentsThatApplyToAResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyAssignments().listForResource("TestResourceGroup",
            "Microsoft.Compute", "virtualMachines/MyTestVm", "domainNames", "MyTestComputer.cloudapp.net", null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
