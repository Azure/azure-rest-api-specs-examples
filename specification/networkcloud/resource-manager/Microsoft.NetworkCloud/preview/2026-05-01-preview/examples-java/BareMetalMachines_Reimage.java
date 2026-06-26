
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineReimageParameters;
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineReimageSafeguardMode;

/**
 * Samples for BareMetalMachines Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/BareMetalMachines_Reimage.json
     */
    /**
     * Sample code: Reimage bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void reimageBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().reimage("resourceGroupName", "bareMetalMachineName",
            new BareMetalMachineReimageParameters().withSafeguardMode(BareMetalMachineReimageSafeguardMode.ALL),
            com.azure.core.util.Context.NONE);
    }
}
