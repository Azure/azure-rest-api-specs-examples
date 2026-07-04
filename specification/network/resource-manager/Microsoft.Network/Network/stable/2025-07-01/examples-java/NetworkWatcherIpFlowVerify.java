
import com.azure.resourcemanager.network.models.Direction;
import com.azure.resourcemanager.network.models.IpFlowProtocol;
import com.azure.resourcemanager.network.models.VerificationIpFlowParameters;

/**
 * Samples for NetworkWatchers VerifyIpFlow.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkWatcherIpFlowVerify.json
     */
    /**
     * Sample code: Ip flow verify.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void ipFlowVerify(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().verifyIpFlow("rg1", "nw1",
            new VerificationIpFlowParameters().withTargetResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1")
                .withDirection(Direction.OUTBOUND).withProtocol(IpFlowProtocol.TCP).withLocalPort("80")
                .withRemotePort("80").withLocalIpAddress("10.2.0.4").withRemoteIpAddress("121.10.1.1"),
            com.azure.core.util.Context.NONE);
    }
}
