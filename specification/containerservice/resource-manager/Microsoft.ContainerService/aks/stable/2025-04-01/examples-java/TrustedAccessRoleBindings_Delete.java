
/**
 * Samples for TrustedAccessRoleBindings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-04-01/examples/
     * TrustedAccessRoleBindings_Delete.json
     */
    /**
     * Sample code: Delete a trusted access role binding.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteATrustedAccessRoleBinding(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getTrustedAccessRoleBindings().delete("rg1",
            "clustername1", "binding1", com.azure.core.util.Context.NONE);
    }
}
