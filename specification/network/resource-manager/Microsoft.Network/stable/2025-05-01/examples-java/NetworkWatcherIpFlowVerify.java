
import com.azure.resourcemanager.network.models.Direction;
import com.azure.resourcemanager.network.models.IpFlowProtocol;
import com.azure.resourcemanager.network.models.VerificationIpFlowParameters;

/**
 * Samples for NetworkWatchers VerifyIpFlow.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherIpFlowVerify.
     * json
     */
    /**
     * Sample code: Ip flow verify.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ipFlowVerify(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkWatchers().verifyIpFlow("rg1", "nw1",
            new VerificationIpFlowParameters()
                .withTargetResourceId(
                    "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1")
                .withDirection(Direction.OUTBOUND).withProtocol(IpFlowProtocol.TCP).withLocalPort("80")
                .withRemotePort("80").withLocalIpAddress("10.2.0.4").withRemoteIpAddress("121.10.1.1"),
            com.azure.core.util.Context.NONE);
    }
}
