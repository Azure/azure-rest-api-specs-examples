
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineCommandSpecification;
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineRunReadCommandsParameters;
import java.util.Arrays;

/**
 * Samples for BareMetalMachines RunReadCommands.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachines_RunReadCommands_Hostname.json
     */
    /**
     * Sample code: Run and retrieve output from a single read only command on bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void runAndRetrieveOutputFromASingleReadOnlyCommandOnBareMetalMachine(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().runReadCommands("resourceGroupName", "bareMetalMachineName",
            new BareMetalMachineRunReadCommandsParameters()
                .withCommands(Arrays.asList(new BareMetalMachineCommandSpecification().withCommand("hostname")))
                .withLimitTimeSeconds(60L),
            com.azure.core.util.Context.NONE);
    }
}
