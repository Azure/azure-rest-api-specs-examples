
/** Samples for Endpoints Start. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Endpoints_Start.json
     */
    /**
     * Sample code: Endpoints_Start.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsStart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getEndpoints().start("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
