import com.azure.core.util.Context;

/** Samples for AfdEndpoints ListByProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_ListByProfile.json
     */
    /**
     * Sample code: AFDEndpoints_ListByProfile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDEndpointsListByProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdEndpoints().listByProfile("RG", "profile1", Context.NONE);
    }
}
