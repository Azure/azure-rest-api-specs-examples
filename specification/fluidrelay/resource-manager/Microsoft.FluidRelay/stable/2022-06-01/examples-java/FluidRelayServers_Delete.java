import com.azure.core.util.Context;

/** Samples for FluidRelayServers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_Delete.json
     */
    /**
     * Sample code: Delete a Fluid Relay server.
     *
     * @param manager Entry point to FluidRelayManager.
     */
    public static void deleteAFluidRelayServer(com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager.fluidRelayServers().deleteWithResponse("myResourceGroup", "myFluidRelayServer", Context.NONE);
    }
}
