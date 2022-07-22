import com.azure.core.util.Context;
import com.azure.resourcemanager.fluidrelay.models.KeyName;
import com.azure.resourcemanager.fluidrelay.models.RegenerateKeyRequest;

/** Samples for FluidRelayServers RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_RegenerateKeys.json
     */
    /**
     * Sample code: Regenerate keys for a Fluid Relay server.
     *
     * @param manager Entry point to FluidRelayManager.
     */
    public static void regenerateKeysForAFluidRelayServer(
        com.azure.resourcemanager.fluidrelay.FluidRelayManager manager) {
        manager
            .fluidRelayServers()
            .regenerateKeyWithResponse(
                "myResourceGroup",
                "myFluidRelayServer",
                new RegenerateKeyRequest().withKeyName(KeyName.KEY1),
                Context.NONE);
    }
}
