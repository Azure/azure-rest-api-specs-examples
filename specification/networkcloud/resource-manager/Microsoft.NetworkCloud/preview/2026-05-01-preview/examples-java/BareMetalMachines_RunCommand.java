
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineRunCommandParameters;
import java.util.Arrays;

/**
 * Samples for BareMetalMachines RunCommand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachines_RunCommand.json
     */
    /**
     * Sample code: Run command on bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        runCommandOnBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().runCommand("resourceGroupName", "bareMetalMachineName",
            new BareMetalMachineRunCommandParameters().withArguments(Arrays.asList("--argument1", "argument2"))
                .withLimitTimeSeconds(60L).withScript("cHdkCg=="),
            com.azure.core.util.Context.NONE);
    }
}
