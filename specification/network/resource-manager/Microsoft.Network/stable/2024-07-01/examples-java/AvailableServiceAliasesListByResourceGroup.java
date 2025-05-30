
/**
 * Samples for AvailableServiceAliases ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * AvailableServiceAliasesListByResourceGroup.json
     */
    /**
     * Sample code: Get available service aliases in the resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAvailableServiceAliasesInTheResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAvailableServiceAliases().listByResourceGroup("rg1",
            "westcentralus", com.azure.core.util.Context.NONE);
    }
}
