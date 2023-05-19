/** Samples for Consoles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Consoles_Delete.json
     */
    /**
     * Sample code: Delete virtual machine console.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteVirtualMachineConsole(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .consoles()
            .delete("resourceGroupName", "virtualMachineName", "default", com.azure.core.util.Context.NONE);
    }
}
