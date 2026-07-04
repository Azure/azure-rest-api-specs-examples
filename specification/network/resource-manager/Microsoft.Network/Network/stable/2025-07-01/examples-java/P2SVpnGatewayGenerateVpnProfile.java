
import com.azure.resourcemanager.network.models.AuthenticationMethod;
import com.azure.resourcemanager.network.models.P2SVpnProfileParameters;

/**
 * Samples for P2SVpnGateways GenerateVpnProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/P2SVpnGatewayGenerateVpnProfile.json
     */
    /**
     * Sample code: GenerateP2SVpnGatewayVPNProfile.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void generateP2SVpnGatewayVPNProfile(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getP2SVpnGateways().generateVpnProfile("rg1", "p2sVpnGateway1",
            new P2SVpnProfileParameters().withAuthenticationMethod(AuthenticationMethod.EAPTLS),
            com.azure.core.util.Context.NONE);
    }
}
