
/**
 * Samples for TrustedAccessRoleBindings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * TrustedAccessRoleBindings_List.json
     */
    /**
     * Sample code: List trusted access role bindings.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTrustedAccessRoleBindings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getTrustedAccessRoleBindings().list("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
