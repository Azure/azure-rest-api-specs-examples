import com.azure.core.util.Context;

/** Samples for FluidRelayServers ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_ListKeys.json
     */
    /**
     * Sample code: Get keys for a Fluid Relay server.
     *
     * @param manager Entry point to FluidRelayManager.
     */
    public static void getKeysForAFluidRelayServer(com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager.fluidRelayServers().listKeysWithResponse("myResourceGroup", "myFluidRelayServer", Context.NONE);
    }
}
