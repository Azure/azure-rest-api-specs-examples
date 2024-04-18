
/**
 * Samples for ResourceUsage List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/ResourceUsage_List.json
     */
    /**
     * Sample code: ResourceUsage_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resourceUsageList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getResourceUsages().list(com.azure.core.util.Context.NONE);
    }
}
