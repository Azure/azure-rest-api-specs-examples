import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.Direction;
import com.azure.resourcemanager.network.models.NetworkConfigurationDiagnosticParameters;
import com.azure.resourcemanager.network.models.NetworkConfigurationDiagnosticProfile;
import java.util.Arrays;

/** Samples for NetworkWatchers GetNetworkConfigurationDiagnostic. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkWatcherNetworkConfigurationDiagnostic.json
     */
    /**
     * Sample code: Network configuration diagnostic.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void networkConfigurationDiagnostic(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkWatchers()
            .getNetworkConfigurationDiagnostic(
                "rg1",
                "nw1",
                new NetworkConfigurationDiagnosticParameters()
                    .withTargetResourceId(
                        "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1")
                    .withProfiles(
                        Arrays
                            .asList(
                                new NetworkConfigurationDiagnosticProfile()
                                    .withDirection(Direction.INBOUND)
                                    .withProtocol("TCP")
                                    .withSource("10.1.0.4")
                                    .withDestination("12.11.12.14")
                                    .withDestinationPort("12100"))),
                Context.NONE);
    }
}
