
/** Samples for Endpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Endpoints_Delete.json
     */
    /**
     * Sample code: Endpoints_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getEndpoints().delete("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
