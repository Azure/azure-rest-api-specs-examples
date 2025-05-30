
import com.azure.resourcemanager.network.models.AzureFirewallNetworkRuleProtocol;
import com.azure.resourcemanager.network.models.AzureFirewallPacketCaptureFlags;
import com.azure.resourcemanager.network.models.AzureFirewallPacketCaptureFlagsType;
import com.azure.resourcemanager.network.models.AzureFirewallPacketCaptureRule;
import com.azure.resourcemanager.network.models.FirewallPacketCaptureParameters;
import java.util.Arrays;

/**
 * Samples for AzureFirewalls PacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/AzureFirewallPacketCapture.
     * json
     */
    /**
     * Sample code: AzureFirewallPacketCapture.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void azureFirewallPacketCapture(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().packetCapture("rg1", "azureFirewall1",
            new FirewallPacketCaptureParameters().withDurationInSeconds(300).withNumberOfPacketsToCapture(5000)
                .withSasUrl("someSASURL").withFileName("azureFirewallPacketCapture")
                .withProtocol(AzureFirewallNetworkRuleProtocol.ANY)
                .withFlags(Arrays.asList(
                    new AzureFirewallPacketCaptureFlags().withType(AzureFirewallPacketCaptureFlagsType.SYN),
                    new AzureFirewallPacketCaptureFlags().withType(AzureFirewallPacketCaptureFlagsType.FIN)))
                .withFilters(Arrays.asList(
                    new AzureFirewallPacketCaptureRule().withSources(Arrays.asList("20.1.1.0"))
                        .withDestinations(Arrays.asList("20.1.2.0")).withDestinationPorts(Arrays.asList("4500")),
                    new AzureFirewallPacketCaptureRule().withSources(Arrays.asList("10.1.1.0", "10.1.1.1"))
                        .withDestinations(Arrays.asList("10.1.2.0")).withDestinationPorts(Arrays.asList("123", "80")))),
            com.azure.core.util.Context.NONE);
    }
}
