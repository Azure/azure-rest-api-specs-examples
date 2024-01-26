
/** Samples for PolicyDefinitions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/
     * deletePolicyDefinition.json
     */
    /**
     * Sample code: Delete a policy definition.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPolicyDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicyDefinitions().deleteWithResponse("ResourceNaming",
            com.azure.core.util.Context.NONE);
    }
}
