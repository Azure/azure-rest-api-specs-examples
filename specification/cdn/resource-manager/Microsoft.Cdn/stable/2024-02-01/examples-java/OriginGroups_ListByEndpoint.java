
/**
 * Samples for OriginGroups ListByEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/OriginGroups_ListByEndpoint.json
     */
    /**
     * Sample code: OriginsGroups_ListByEndpoint.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originsGroupsListByEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getOriginGroups().listByEndpoint("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
