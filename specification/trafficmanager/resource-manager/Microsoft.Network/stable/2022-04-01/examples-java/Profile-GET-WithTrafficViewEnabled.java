
/**
 * Samples for Profiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-GET-
     * WithTrafficViewEnabled.json
     */
    /**
     * Sample code: Profile-GET-WithTrafficViewEnabled.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profileGETWithTrafficViewEnabled(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getProfiles().getByResourceGroupWithResponse(
            "azuresdkfornetautoresttrafficmanager1323", "azuresdkfornetautoresttrafficmanager3880",
            com.azure.core.util.Context.NONE);
    }
}
