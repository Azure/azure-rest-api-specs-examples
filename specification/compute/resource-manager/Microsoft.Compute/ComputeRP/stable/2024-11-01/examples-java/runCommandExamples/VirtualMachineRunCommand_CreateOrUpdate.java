
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineRunCommandInner;
import com.azure.resourcemanager.compute.models.RunCommandInputParameter;
import com.azure.resourcemanager.compute.models.RunCommandManagedIdentity;
import com.azure.resourcemanager.compute.models.VirtualMachineRunCommandScriptSource;
import java.util.Arrays;

/**
 * Samples for VirtualMachineRunCommands CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/runCommandExamples/
     * VirtualMachineRunCommand_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a run command.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateARunCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineRunCommands().createOrUpdate(
            "myResourceGroup", "myVM", "myRunCommand",
            new VirtualMachineRunCommandInner().withLocation("West US")
                .withSource(new VirtualMachineRunCommandScriptSource()
                    .withScriptUri("https://mystorageaccount.blob.core.windows.net/scriptcontainer/scriptURI"))
                .withParameters(Arrays.asList(new RunCommandInputParameter().withName("param1").withValue("value1"),
                    new RunCommandInputParameter().withName("param2").withValue("value2")))
                .withAsyncExecution(false).withRunAsUser("user1").withRunAsPassword("fakeTokenPlaceholder")
                .withTimeoutInSeconds(3600)
                .withOutputBlobUri(
                    "https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt")
                .withErrorBlobUri("https://mystorageaccount.blob.core.windows.net/scriptcontainer/scriptURI")
                .withOutputBlobManagedIdentity(
                    new RunCommandManagedIdentity().withClientId("22d35efb-0c99-4041-8c5b-6d24db33a69a"))
                .withTreatFailureAsDeploymentFailure(false),
            com.azure.core.util.Context.NONE);
    }
}
