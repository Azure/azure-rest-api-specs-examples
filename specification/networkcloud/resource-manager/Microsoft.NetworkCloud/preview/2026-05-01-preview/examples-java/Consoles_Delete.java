
/**
 * Samples for Consoles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Consoles_Delete.json
     */
    /**
     * Sample code: Delete virtual machine console.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteVirtualMachineConsole(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.consoles().delete("resourceGroupName", "virtualMachineName", "default", null, null,
            com.azure.core.util.Context.NONE);
    }
}
