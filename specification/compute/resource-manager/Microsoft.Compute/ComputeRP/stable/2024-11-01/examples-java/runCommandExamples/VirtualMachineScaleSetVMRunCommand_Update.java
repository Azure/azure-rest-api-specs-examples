
import com.azure.resourcemanager.compute.models.RunCommandManagedIdentity;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandScriptSource;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandUpdate;

/**
 * Samples for VirtualMachineScaleSetVMRunCommands Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/runCommandExamples/
     * VirtualMachineScaleSetVMRunCommand_Update.json
     */
    /**
     * Sample code: Update VirtualMachineScaleSet VM run command.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateVirtualMachineScaleSetVMRunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMRunCommands().update(
            "myResourceGroup", "myvmScaleSet", "0", "myRunCommand",
            new VirtualMachineRunCommandUpdate().withSource(new VirtualMachineRunCommandScriptSource()
                .withScriptUri("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1")
                .withScriptUriManagedIdentity(
                    new RunCommandManagedIdentity().withObjectId("4231e4d2-33e4-4e23-96b2-17888afa6072"))),
            com.azure.core.util.Context.NONE);
    }
}
