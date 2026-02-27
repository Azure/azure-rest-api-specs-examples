
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/UsageList.json
     */
    /**
     * Sample code: List usages.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getUsages().list("westus", com.azure.core.util.Context.NONE);
    }
}
