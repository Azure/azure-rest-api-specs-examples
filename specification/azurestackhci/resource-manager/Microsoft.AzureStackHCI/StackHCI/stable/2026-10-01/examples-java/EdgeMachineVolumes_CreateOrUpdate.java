
import com.azure.resourcemanager.azurestackhci.models.EdgeMachineVolumeProperties;
import com.azure.resourcemanager.azurestackhci.models.VolumeConfiguration;

/**
 * Samples for EdgeMachineVolumes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeMachineVolumes_CreateOrUpdate.json
     */
    /**
     * Sample code: EdgeMachineVolumes_CreateOrUpdate.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineVolumesCreateOrUpdate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineVolumes().define("vol-001").withExistingEdgeMachine("test-rg", "EdgeMachine01")
            .withProperties(
                new EdgeMachineVolumeProperties().withVolumeConfiguration(new VolumeConfiguration().withReserved("")))
            .create();
    }
}
