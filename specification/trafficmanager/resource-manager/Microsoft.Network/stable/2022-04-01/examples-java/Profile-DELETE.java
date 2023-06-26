/** Samples for Profiles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-DELETE.json
     */
    /**
     * Sample code: Profile-DELETE.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profileDELETE(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .trafficManagerProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .deleteWithResponse(
                "azuresdkfornetautoresttrafficmanager1323",
                "azuresdkfornetautoresttrafficmanager3880",
                com.azure.core.util.Context.NONE);
    }
}
