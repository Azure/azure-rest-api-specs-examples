import com.azure.core.util.Context;

/** Samples for PolicySetDefinitions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicySetDefinitions.json
     */
    /**
     * Sample code: List policy set definitions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPolicySetDefinitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getPolicySetDefinitions().list(null, null, Context.NONE);
    }
}
