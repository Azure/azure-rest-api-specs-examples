import com.azure.core.util.Context;

/** Samples for FluidRelayContainers ListByFluidRelayServers. */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_ListByFluidRelayServer.json
     */
    /**
     * Sample code: List all Fluid Relay containers in a Fluid Relay server.
     *
     * @param manager Entry point to FluidRelayManager.
     */
    public static void listAllFluidRelayContainersInAFluidRelayServer(
        com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager.fluidRelayContainers().listByFluidRelayServers("myResourceGroup", "myFluidRelayServer", Context.NONE);
    }
}
