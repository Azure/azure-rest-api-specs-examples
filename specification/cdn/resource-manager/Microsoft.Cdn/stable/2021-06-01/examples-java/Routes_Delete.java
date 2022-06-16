import com.azure.core.util.Context;

/** Samples for Routes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Routes_Delete.json
     */
    /**
     * Sample code: Routes_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routesDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getRoutes()
            .delete("RG", "profile1", "endpoint1", "route1", Context.NONE);
    }
}
