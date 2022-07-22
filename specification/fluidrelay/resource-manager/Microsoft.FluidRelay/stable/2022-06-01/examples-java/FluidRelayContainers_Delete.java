import com.azure.core.util.Context;

/** Samples for FluidRelayContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_Delete.json
     */
    /**
     * Sample code: Delete a Fluid Relay container.
     *
     * @param manager Entry point to FluidRelayManager.
     */
    public static void deleteAFluidRelayContainer(com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager
            .fluidRelayContainers()
            .deleteWithResponse("myResourceGroup", "myFluidRelayServer", "myFluidRelayContainer", Context.NONE);
    }
}
