
import com.azure.resourcemanager.networkcloud.models.AdministrativeCredentials;
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineReplaceParameters;
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineReplaceSafeguardMode;
import com.azure.resourcemanager.networkcloud.models.BareMetalMachineReplaceStoragePolicy;

/**
 * Samples for BareMetalMachines Replace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * BareMetalMachines_Replace.json
     */
    /**
     * Sample code: Replace bare metal machine.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void replaceBareMetalMachine(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bareMetalMachines().replace("resourceGroupName", "bareMetalMachineName",
            new BareMetalMachineReplaceParameters()
                .withBmcCredentials(
                    new AdministrativeCredentials().withPassword("fakeTokenPlaceholder").withUsername("bmcuser"))
                .withBmcMacAddress("00:00:4f:00:57:ad").withBootMacAddress("00:00:4e:00:58:af").withMachineName("name")
                .withSafeguardMode(BareMetalMachineReplaceSafeguardMode.ALL).withSerialNumber("BM1219XXX")
                .withStoragePolicy(BareMetalMachineReplaceStoragePolicy.DISCARD_ALL),
            com.azure.core.util.Context.NONE);
    }
}
