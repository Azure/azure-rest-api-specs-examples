
import com.azure.resourcemanager.network.fluent.models.NspProfileInner;

/**
 * Samples for NetworkSecurityPerimeterProfiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspProfilePut.json
     */
    /**
     * Sample code: NspProfilesPut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspProfilesPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterProfiles().createOrUpdateWithResponse(
            "rg1", "nsp1", "profile1", new NspProfileInner(), com.azure.core.util.Context.NONE);
    }
}
