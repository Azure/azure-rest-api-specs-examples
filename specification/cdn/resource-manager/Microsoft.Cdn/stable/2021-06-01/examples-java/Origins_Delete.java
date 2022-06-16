import com.azure.core.util.Context;

/** Samples for Origins Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Delete.json
     */
    /**
     * Sample code: Origins_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getOrigins()
            .delete("RG", "profile1", "endpoint1", "origin1", Context.NONE);
    }
}
