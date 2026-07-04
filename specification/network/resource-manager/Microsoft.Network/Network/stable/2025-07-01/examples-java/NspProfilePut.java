
import com.azure.resourcemanager.network.fluent.models.NspProfileInner;

/**
 * Samples for NetworkSecurityPerimeterProfiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspProfilePut.json
     */
    /**
     * Sample code: NspProfilesPut.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspProfilesPut(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterProfiles().createOrUpdateWithResponse("rg1", "nsp1",
            "profile1", new NspProfileInner(), com.azure.core.util.Context.NONE);
    }
}
