
/** Samples for PolicyDefinitions Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinition.
     * json
     */
    /**
     * Sample code: Retrieve a policy definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAPolicyDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitions().getWithResponse("ResourceNaming",
            com.azure.core.util.Context.NONE);
    }
}
