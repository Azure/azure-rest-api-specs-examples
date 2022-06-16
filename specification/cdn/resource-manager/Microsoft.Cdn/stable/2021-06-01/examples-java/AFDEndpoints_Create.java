import com.azure.core.util.Context;

/** Samples for AfdEndpoints Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_Create.json
     */
    /**
     * Sample code: AFDEndpoints_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDEndpointsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdEndpoints()
            .create("RG", "profile1", "endpoint1", null, Context.NONE);
    }
}
