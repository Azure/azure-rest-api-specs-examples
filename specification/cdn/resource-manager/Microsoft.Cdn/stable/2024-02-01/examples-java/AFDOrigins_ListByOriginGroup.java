
/**
 * Samples for AfdOrigins ListByOriginGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOrigins_ListByOriginGroup.json
     */
    /**
     * Sample code: AFDOrigins_ListByOriginGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginsListByOriginGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdOrigins().listByOriginGroup("RG", "profile1",
            "origingroup1", com.azure.core.util.Context.NONE);
    }
}
