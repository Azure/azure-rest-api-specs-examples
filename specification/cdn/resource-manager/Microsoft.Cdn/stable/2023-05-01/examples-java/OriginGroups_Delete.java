
/** Samples for OriginGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/OriginGroups_Delete.json
     */
    /**
     * Sample code: OriginGroups_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originGroupsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getOriginGroups().delete("RG", "profile1", "endpoint1",
            "originGroup1", com.azure.core.util.Context.NONE);
    }
}
