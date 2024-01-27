
/** Samples for Profiles ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/trafficmanager/resource-manager/Microsoft.Network/stable/2022-04-01/examples/Profile-GET-
     * ByResourceGroup.json
     */
    /**
     * Sample code: ListProfilesByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProfilesByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getProfiles()
            .listByResourceGroup("azuresdkfornetautoresttrafficmanager3640", com.azure.core.util.Context.NONE);
    }
}
