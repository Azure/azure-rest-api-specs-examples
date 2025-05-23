
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineRunCommandInner;
import com.azure.resourcemanager.compute.models.RunCommandInputParameter;
import com.azure.resourcemanager.compute.models.RunCommandManagedIdentity;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandScriptSource;
import java.util.Arrays;

/**
 * Samples for VirtualMachineScaleSetVMRunCommands CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/runCommandExamples/
     * VirtualMachineScaleSetVMRunCommand_CreateOrUpdate.json
     */
    /**
     * Sample code: Create VirtualMachineScaleSet VM run command.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualMachineScaleSetVMRunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMRunCommands().createOrUpdate(
            "myResourceGroup", "myvmScaleSet", "0", "myRunCommand",
            new VirtualMachineRunCommandInner().withLocation("West US")
                .withSource(new VirtualMachineRunCommandScriptSource()
                    .withScriptUri("https://mystorageaccount.blob.core.windows.net/scriptcontainer/MyScript.ps1")
                    .withScriptUriManagedIdentity(
                        new RunCommandManagedIdentity().withObjectId("4231e4d2-33e4-4e23-96b2-17888afa6072")))
                .withParameters(Arrays.asList(new RunCommandInputParameter().withName("param1").withValue("value1"),
                    new RunCommandInputParameter().withName("param2").withValue("value2")))
                .withAsyncExecution(false).withRunAsUser("user1").withRunAsPassword("fakeTokenPlaceholder")
                .withTimeoutInSeconds(3600)
                .withOutputBlobUri(
                    "https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt")
                .withErrorBlobUri("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt")
                .withOutputBlobManagedIdentity(
                    new RunCommandManagedIdentity().withClientId("22d35efb-0c99-4041-8c5b-6d24db33a69a"))
                .withErrorBlobManagedIdentity(new RunCommandManagedIdentity())
                .withTreatFailureAsDeploymentFailure(true),
            com.azure.core.util.Context.NONE);
    }
}
