
/**
 * Samples for TrustedAccessRoles List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * TrustedAccessRoles_List.json
     */
    /**
     * Sample code: List trusted access roles.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTrustedAccessRoles(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getTrustedAccessRoles().list("westus2",
            com.azure.core.util.Context.NONE);
    }
}
